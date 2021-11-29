package token

import (
    "github.com/joho/godotenv"
    "os"
    "log"
    "fmt"
    "regexp"
    "strings"
    "bytes"
    "strconv"
    "database/sql"
    "github.com/go-sql-driver/mysql"
    "math/rand"
    "math"
    "encoding/json"
)

var wordRe *regexp.Regexp 
var nonLetterRe *regexp.Regexp 
var wsRe *regexp.Regexp 

var db *sql.DB 

var ChainIds []uint32

func normalizeAddressString(s string) string {
    s = strings.ToLower(strings.TrimSpace(s))
    if len(s) > 1 && s[0] == 48 && s[1] == 120 {
        s = s[2:]
    }
    return s
}

func stringHash(s string) int32 {
    var h int32 = 5381
    for _, r := range s {
        h = ((h << 5) + h) + int32(r)
    }
    return h
}

func addWordHashes(w string, p *string, hmap map[int32]int32, recurse bool) {
    var hash int32 = 0
    var weight int32 = 0
    w = strings.ToLower(w)
    if recurse {
        var pp string = ""
        for _, ww := range nonLetterRe.Split(w, -1) {
            addWordHashes(ww, &pp, hmap, false)
        }
    }
    if t := nonLetterRe.ReplaceAllString(w, ""); len(t) > 0 {
        w = t
    }
    hash = stringHash(w)
    weight = int32(len(w))
    if weight > 2 {
        hmap[hash] += weight
    }
    if len(*p) > 1 && len(w) > 1 {
        lw := *p + w
        hash = stringHash(lw)
        weight = int32(len(lw) / 2)
        hmap[hash] += weight
    }
    if len(w) > 0 {
        *p = w
    }
}

func wordHashes(s string) map[int32]int32 {
    splitted := wsRe.Split(strings.TrimSpace(s), -1)
    var p string = ""
    var pp string = ""
    hmap := make(map[int32]int32)
    for _, w := range splitted {
        for _, ww := range wordRe.FindAllString(w, -1) {
            addWordHashes(ww, &pp, hmap, true)
        }
        addWordHashes(w, &p, hmap, true)
    }
    return hmap
}


type Token struct {
    Id uint64 
    Name string 
    Minter string 
    Veracity int32
    Data string
}

func PutToken(t Token, chainId uint32) {

    name := strings.TrimSpace(t.Name)
    _, insertErr := db.Exec(
        `INSERT INTO tokens 
        (chain_id, token_id, name_hash, data, minter, veracity)
        VALUES (?, ?, ?, ?, ?, ?)`, 
        chainId, t.Id, stringHash(name), t.Data, 
        normalizeAddressString(t.Minter), t.Veracity)

    if insertErr != nil {
        queryError, _ := insertErr.(*mysql.MySQLError)
        if queryError.Number == 1062 {
            // We need to use exec whenever we do insert, update, delete
            // to prevent leaking the connections
            db.Exec(`UPDATE tokens SET veracity = ?, data = ? 
                WHERE chain_id = ? AND token_id = ?`, 
                t.Veracity, t.Data, chainId, t.Id)
        }

    } else {
        for hash, weight := range wordHashes(name) {
            db.Exec(`INSERT INTO lookup 
                (word_hash, weight, chain_id, token_id) 
                VALUES (?, ?, ?, ?)`, hash, weight, chainId, t.Id)
        }        
    }
}

func extractData(t *Token) bool {
    var r []string
    if e := json.Unmarshal([]byte(t.Data), &r); e == nil {
        t.Name = r[1]
        t.Minter = r[6]
        return true
    }
    return false
}

func MissingIds(totalSupply uint64, chainId uint32) ([]uint64) {
     
    missing := make([]uint64, 0, 32)
    rows, e := db.Query(`SELECT token_id FROM tokens 
        WHERE chain_id = ? ORDER BY token_id ASC`, chainId)
    if e != nil {
        return missing
    }

    var i uint64 = 0
    for rows.Next() {
        var tokenId uint64 = 0
        if e := rows.Scan(&tokenId); e == nil {
            for j := i; j < tokenId; j++ {
                missing = append(missing, j)
            }
            i = tokenId + 1
        }
    }

    for j := i; j < totalSupply; j++ {
        missing = append(missing, j)
    }
    rows.Close()
    return missing
}

func SearchName(name string, n uint64, chainId uint32) ([]Token) {
    
    name = strings.TrimSpace(name)
    tokens := make([]Token, 0, n)
    t := Token{}

    var buffer bytes.Buffer

    for hash, _ := range wordHashes(name) {
        if buffer.Len() > 0 {
            buffer.WriteString(",")    
        }
        buffer.WriteString(strconv.FormatInt(int64(hash), 10))
    }

    stm := `SELECT a.token_id, a.data, a.veracity, b.w FROM tokens a 
        INNER JOIN (
            SELECT 
                l.chain_id,
                l.token_id,
                SUM(l.weight) AS w, 
                (SELECT t.veracity FROM tokens t WHERE 
                    t.chain_id = l.chain_id AND t.token_id = l.token_id) AS v 
            FROM lookup l
            WHERE l.word_hash IN (` + buffer.String() + `) 
            AND l.chain_id = ` + strconv.FormatUint(uint64(chainId), 10) + `
            GROUP BY l.token_id
            ORDER BY w DESC, v DESC
            LIMIT ` + strconv.FormatUint(n, 10) + `
        ) AS b 
        ON a.chain_id = b.chain_id AND a.token_id = b.token_id`
    approxRows, e := db.Query(stm)

    if e != nil {
        return tokens
    }
    for approxRows.Next() {
        var weight float64 = 0
        e := approxRows.Scan(&t.Id, &t.Data, &t.Veracity, &weight)
        if e == nil && extractData(&t) {
            tokens = append(tokens, t) 
        } 
    }
    approxRows.Close()
    return tokens
}

func SearchMinter(minter string, pageSize uint64, 
    pageNum uint64, chainId uint32) ([]Token, uint64) {

    tokens := make([]Token, 0, pageSize)
    t := Token{}

    minter = normalizeAddressString(minter)
    var rowCount uint64 = 0
    e := db.QueryRow(`SELECT COUNT(token_id) FROM tokens 
        WHERE minter = ? AND chain_id = ?`, minter, chainId).Scan(&rowCount)
    if e != nil {
        return tokens, 0
    }
    numPages := uint64(math.Ceil(float64(rowCount) / float64(pageSize)))
    if pageNum < 0 {
        pageNum = 0
    } else if pageNum >= numPages {
        pageNum = numPages - 1
    }
    offset := pageNum * pageSize
    rows, e := db.Query(`SELECT token_id, data, veracity 
        FROM tokens WHERE minter = ? AND chain_id = ?
        ORDER BY veracity DESC
        LIMIT ?, ?`, minter, chainId, offset, pageSize)
    if e != nil {
        return tokens, 0   
    }
    for rows.Next() {
        e := rows.Scan(&t.Id, &t.Data, &t.Veracity)
        if e == nil && extractData(&t) {
            tokens = append(tokens, t)
        }
    }
    rows.Close()
    return tokens, numPages
}


func randomString(n int) string {
    var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func randomName() string {
    n := 2 + rand.Intn(10)
    s := ""
    for i := 0; i < n; i++ {
        if i > 0 {
            s = s + " "    
        }
        s = s + randomString(5)
    }
    return s
}

func randomToken(id uint64) Token {
    name := randomName()
    data := `["` + strconv.Itoa(int(id)) + `","` + name + `","","","","",`+
        `"0x69ba5ea684410246068837088c271635e7aa3fbb",` +
        `"0x3039","0x96b441",` +
        `"4524371646796053238226387062396914999321789068020079075569951078729989172462",` +
        `"0","0"]`
    return Token{
        id, 
        name, 
        "0x69ba5ea684410246068837088c271635e7aa3fbb", 
        int32(rand.Intn(10)),
        data,
    }
}

func PopulateRandom() {
    var i uint64
    for i = 0; i < 10000; i++ {
        t := randomToken(i)
        PutToken(t, 4)    
    }
}

func init() {
    wordRe = regexp.MustCompile(`[\x{2190}-\x{10FFFF}]|[\p{Lu}][\p{Ll}\']+|[\d]+`)
    nonLetterRe = regexp.MustCompile(`[^\p{Ll}]`)
    wsRe = regexp.MustCompile(`[\s_\-]+`)

    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
    dbName := os.Getenv("DB_NAME")
    dbPass := os.Getenv("DB_PASS")
    dbUser := os.Getenv("DB_USER")
    dbHost := os.Getenv("DB_HOST")

    t, err := sql.Open("mysql", dbUser + ":" + dbPass + "@" + dbHost + "/" + dbName)
    if err != nil {
        log.Fatal(err)
    }
    db = t
    fmt.Println("Connected to tokens database!")

    ChainIds = make([]uint32, 0)
    for _, c := range strings.Split(os.Getenv("CHAIN_IDS"), ",") {
        if i, err := strconv.ParseUint(c, 10, 64); err == nil {
            ChainIds = append(ChainIds, uint32(i))
        }
    }
}
