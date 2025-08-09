// cmd/server/main.go
package main

import (
	"dfiler/internal/config"
	"flag"
	"fmt"
	"log"
)

func main() {
	configPath := flag.String("config", "etc/config.json", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("Server will start on port %v\n", cfg.ServerPort)
	// TODO: 啟動 server 等
}
