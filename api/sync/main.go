package sync

import (
    "github.com/joho/godotenv"
    "os"

    "fmt"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/chainrand/api/cc"
    "github.com/chainrand/api/token"
    "github.com/ethereum/go-ethereum/crypto"

    "log"
    "context"
    "strings"
    "strconv"
    "math/big" 
    "time"
    "encoding/json"
)

type eventSig struct {
    SigHex string
    Name string
}

func newEventSig(sig string) *eventSig {
    name := ""
    if o := strings.Index(sig, "("); o >= 0 {
        name = sig[:o]
    }
    return &eventSig{crypto.Keccak256Hash([]byte(sig)).Hex(), name}
}

type eventUnpacker struct {
    Abi *abi.ABI
    VLog *types.Log
    SigHex string
}

func newEventUnpacker(abi *abi.ABI, vLog *types.Log) *eventUnpacker {
    return &eventUnpacker{abi, vLog, vLog.Topics[0].Hex()}
}

func (_eu *eventUnpacker) GetTokenId(es *eventSig) (uint64, bool) {
    if _eu.SigHex == es.SigHex {
        unpacked, err := _eu.Abi.Unpack(es.Name, _eu.VLog.Data)
        if err == nil { 
            if b, ok := unpacked[0].(*big.Int); ok {
                return b.Uint64(), true
            }
        }
    }
    return 0, false
}

func veracity(paid string, verified bool) int32 {
    if verified {
        return 2147483647
    }
    p := new(big.Int)
    p.SetString(paid, 10)
    d := new(big.Int)
    d.SetString("1000000000000000", 10)
    q := new(big.Int).Div(p, d)
    m := new(big.Int)
    m.SetString("2147483646", 10)
    if q.Cmp(m) > 0 {
        q = m
    }
    return int32(q.Int64())
}

func downloadTokens(contractClient *cc.Cc, tokenIds []uint64) ([]token.Token, error) {
    
    ids := make([]*big.Int, len(tokenIds))
    for i, t := range tokenIds {
        b := new(big.Int)
        b.SetUint64(t)
        ids[i] = b
    }

    tokens := make([]token.Token, 0, len(tokenIds))
    
    r, err := contractClient.TokenData(&bind.CallOpts{}, ids)

    if err != nil {
        return tokens, err    
    }

    n := int(len(r) / 12)
    for j := 0; j < n; j++ {
        t := token.Token{}    
        o := j * 12
        if tokenId, err := strconv.ParseUint(r[o], 10, 64); err == nil {
            t.Id = tokenId
            t.Name = strings.TrimSpace(r[o+1])
            t.Minter = r[o+6]
            t.Veracity = veracity(r[o+10], r[o+11] == "1")
            b, _ := json.Marshal(r[o:o+12])
            t.Data = string(b)
            tokens = append(tokens, t)
        }
    }
    return tokens, nil
}

func syncToken(contractClient *cc.Cc, tokenId uint64, chainId uint32) {
    tokenIds := []uint64{tokenId}
    for i := 0; i < 3; i++ {
        if tokens, err := downloadTokens(contractClient, tokenIds); err == nil {
            for _, t := range tokens {
                token.PutToken(t, chainId)
                return 
            }
        }
        time.Sleep(3000 * time.Millisecond)
    }
}

func syncMissing(contractClient *cc.Cc, chainId uint32) {
    
    r, err := contractClient.TotalSupply(&bind.CallOpts{})
    
    if err != nil {
        return
    }

    totalSupply := r.Uint64()
    missing := token.MissingIds(totalSupply, chainId)

    step := 10
    n := len(missing)
    for i := 0; i < n; i += step {
        end := i + step
        if end >= n {
            end = n
        }
        tokenIds := missing[i:end]
        if tokens, err := downloadTokens(contractClient, tokenIds); err == nil {
            for _, t := range tokens {
                token.PutToken(t, chainId) 
            }
        }
    }
}

func subscribe(wssEndpoint string, contractAddress string, chainId uint32) error {

    client, err := ethclient.Dial(wssEndpoint)
    if err != nil { 
        log.Fatal(err) 
    }

    pf := "(" + strconv.FormatUint(uint64(chainId), 10) + "): "

    fmt.Println(pf + "Connected to infura!")
    
    parsedAddress := common.HexToAddress(contractAddress)
    
    contractClient, err := cc.NewCc(parsedAddress, client)
    if err != nil { 
        return err
    }

    fmt.Println(pf + "Connected to contract!")


    syncMissing(contractClient, chainId)
    // return nil

    query := ethereum.FilterQuery{
        Addresses: []common.Address{parsedAddress},
    }

    logs := make(chan types.Log)

    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil { 
        return err
    }

    fmt.Println(pf + "Subscribed to events!")

    contractAbi, err := abi.JSON(strings.NewReader(string(cc.CcMetaData.ABI)))
    if err != nil { 
        return err
    }

    eVerified := newEventSig("Verified(uint256)")
    eRandomnessFufilled := newEventSig("RandomnessFufilled(uint256)")
    eSeedKeyRevealed := newEventSig("SeedKeyRevealed(uint256)")
    ePaid := newEventSig("Paid(uint256)")
    
    for {
        select {
        case err := <-sub.Err():
            return err
        case vLog := <-logs:
            eu := newEventUnpacker(&contractAbi, &vLog)
            if tokenId, ok := eu.GetTokenId(eVerified); ok {
                go syncToken(contractClient, tokenId, chainId)
            }
            if tokenId, ok := eu.GetTokenId(eRandomnessFufilled); ok {
                go syncToken(contractClient, tokenId, chainId)   
            }
            if tokenId, ok := eu.GetTokenId(eSeedKeyRevealed); ok {
                go syncToken(contractClient, tokenId, chainId)
            }
            if tokenId, ok := eu.GetTokenId(ePaid); ok {
                go syncToken(contractClient, tokenId, chainId)
            }
        }
    }
}

func subscribeLoop(wssEndpoint string, contractAddress string, chainId uint32) {
    for {
        subscribe(wssEndpoint, contractAddress, chainId)
        // return 
    }
}

func Run() {

    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    for _, c := range token.ChainIds {
        s := strconv.FormatUint(uint64(c), 10)
        go subscribeLoop(
            os.Getenv("WSS_ENDPOINT_" + s), 
            os.Getenv("CONTRACT_ADDRESS_" + s), 
            c)    
    }
    
}