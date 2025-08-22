package main

import (
	"fmt"
	"log"
	"projeto-modelo/configs"
)

func main() {
    config, err := configs.LoadConfig(".")
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }
    fmt.Println(config.DBDriver)
}
