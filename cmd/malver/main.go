package main

import (
	"flag"
	"log"
	"fmt"
	"os"

	"github.com/omarelshopky/malver/config"
	"github.com/omarelshopky/malver/internal/handlers"
	"github.com/omarelshopky/malver/internal/server"
)

func main() {
	cfg := config.LoadConfig()

	flag.Parse()

	if cfg.PrintUploadCommands {
		if cfg.AttackingIP == "<ATTACKING_IP>" || cfg.FilePath == "<FILE_PATH>" {
			fmt.Println("[WARNING] Both -ip and -file should be used with -upload-commands to generate ready-to-use upload commands.")
		}

		handlers.GenerateUploadCommands(cfg.UploadEndpoint, cfg.AttackingIP, cfg.Port, cfg.FilePath)

		return
	}

	if err := os.MkdirAll(cfg.UploadDir, 0755); err != nil {
		log.Fatalf("Failed to create upload directory: %v", err)
	}

	if err := os.MkdirAll(cfg.DownloadDir, 0755); err != nil {
		log.Fatalf("Failed to create download directory: %v", err)
	}

	server.Start(cfg)
}