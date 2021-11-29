package server

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/chainrand/api/token"
    "math/big"
    "strings"
    "strconv"    
)

var urlPathRoot string = "chainrand/api"

type ResponseToken struct {
    Id   uint64 `json:"id"`
    Name string `json:"name"`
    CodeURI string `json:"codeURI"`
    SeedKey string `json:"seedKey"`
    ImageURI string `json:"imageURI"`
    ProjectURI string `json:"projectURI"`
    Minter string `json:"minter"`
    SeedKeyHash string `json:"seedKeyHash"`
    CodeHash string `json:"codeHash"`
    Randomness string `json:"randomness"`
    Paid string `json:"paid"`
    Verified bool `json:"verified"`
}

type SearchNameResponse struct {
    Tokens []ResponseToken `json:"tokens"`
}

type SearchMinterResponse struct {
    Tokens []ResponseToken `json:"tokens"`
    PageCount uint64 `json:"pageCount"`
}

func clamp(x uint64, low uint64, high uint64) uint64 { 
    if high < low {
        low, high = high, low
    }
    if x > high { 
        x = high
    }
    if x < low {
        x = low
    }
    return x
}

func formUint64(r *http.Request, key string, defaultValue uint64) uint64 {
    if x := r.FormValue(key); len(x) > 0 {
        b := new(big.Int)
        b.SetString(x, 10)
        if b.IsUint64() {
            return b.Uint64()
        }
    }
    return defaultValue
}

func ParseToken(t token.Token) (ResponseToken, bool) {
    var r []string
    u := ResponseToken{}
    if err := json.Unmarshal([]byte(t.Data), &r); err == nil {
        u.Id = t.Id
        u.Name = r[1]
        u.CodeURI = r[2]
        u.SeedKey = r[3]
        u.ImageURI = r[4]
        u.ProjectURI = r[5]
        u.Minter = r[6]
        u.SeedKeyHash = r[7]
        u.CodeHash = r[8]
        u.Randomness = r[9]
        u.Paid = r[10]
        u.Verified = r[11] == "1"
        return u, true
    }
    return u, false
}

func outputJsonString(w http.ResponseWriter, s []byte) {
    w.Header().Set("Content-type", "application/json; charset=utf-8")
    fmt.Fprintf(w, "%s", s)
}

func parseChainId(c string) (uint32, bool) {
    c = strings.TrimSpace(c)
    for _, x := range token.ChainIds {
        if strconv.FormatUint(uint64(x), 10) == c {
            return x, true
        }
    }
    return 0, false
}

func RunServer(w http.ResponseWriter, r *http.Request) {
     
    if origin := r.Header.Get("Origin"); origin != "" {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers",
            "X-Requested-With, Content-Type, Origin, Cache-Control, " +
            "Pragma, Authorization, Accept, Accept-Encoding")
    }

    if r.Method == "OPTIONS" {
        return
    }
    
    if err := r.ParseForm() ; err != nil {
        return
    }

    if chainId, ok := parseChainId(r.FormValue("c")); ok {

        if name := r.FormValue("name"); len(name) > 0 {
            numSimilar := clamp(formUint64(r, "n", 10), 10, 100)
            l := SearchNameResponse{make([]ResponseToken, 0)}
            tokens := token.SearchName(
                name, numSimilar, chainId)
            for _, t := range tokens {
                if u, ok := ParseToken(t); ok {
                    l.Tokens = append(l.Tokens, u)    
                }
            }
            if ms, err := json.Marshal(l); err == nil {
                outputJsonString(w, ms)
            }
            return
        }

        if minter := r.FormValue("minter"); len(minter) > 0 {
            pageSize := clamp(formUint64(r, "n", 10), 10, 100)
            pageNum := formUint64(r, "p", 0)
            tokens, pageCount := token.SearchMinter(
                minter, pageSize, pageNum, chainId)
            l := SearchMinterResponse{make([]ResponseToken, 0), pageCount}
            for _, t := range tokens {
                if u, ok := ParseToken(t); ok {
                    l.Tokens = append(l.Tokens, u)    
                }
            }
            if ms, err := json.Marshal(l); err == nil {
                outputJsonString(w, ms)
            }
            return
        }    
    }
}

func Run() {
    
    fmt.Println("Running server!")
    http.HandleFunc("/", RunServer)
    http.ListenAndServe(":8081", nil)
}

