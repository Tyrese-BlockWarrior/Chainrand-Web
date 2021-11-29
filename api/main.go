package main

import (
    "github.com/chainrand/api/server"
    "github.com/chainrand/api/sync"
)

func main() {
    sync.Run()
    server.Run()
}