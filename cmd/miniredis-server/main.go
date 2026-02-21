package main

import (
	"log"

	"github.com/koraydinc/mini-redis/internal/config"
	"github.com/koraydinc/mini-redis/internal/server"
)

func main() {
	cfg := config.Load()
	tcpServer := server.NewTCPServer(cfg.Addr)

	if err := tcpServer.Start(); err != nil {
		log.Fatal(err)
	}
}
