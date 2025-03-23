package main

import (
	"log"
	"os"
	"github.com/omarelshopky/malver/internal/server"
	"github.com/omarelshopky/malver/config"
)

func main() {
	cfg := config.LoadConfig()

	if err := os.MkdirAll(cfg.UploadDir, 0755); err != nil {
		log.Fatalf("Failed to create upload directory: %v", err)
	}

	if err := os.MkdirAll(cfg.DownloadDir, 0755); err != nil {
		log.Fatalf("Failed to create download directory: %v", err)
	}

	server.Start(cfg)
}