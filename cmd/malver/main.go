package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/omarelshopky/malver/config"
	"github.com/omarelshopky/malver/internal/handlers"
	"github.com/omarelshopky/malver/internal/logger"
	"github.com/omarelshopky/malver/internal/server"
)

func main() {
	cfg := config.LoadConfig()

	if cfg.PrintUploadCommands.Enabled {
		printUploadCommands(cfg.PrintUploadCommands, cfg.Server.Port, cfg.Endpoints.Upload)

		return
	}

	createApplicationDirs(cfg.Server.Dirs)

	logger.InitLogger(&cfg.Logging)

	server.Start(cfg.Server.Port, cfg.Endpoints, cfg.Server.Dirs)
}

func printUploadCommands(cfg config.PrintUploadCommandsConfig, AttackingPort, UploadEndpoint string) {
	if cfg.AttackingIP == "<ATTACKING_IP>" || cfg.FilePath == "<FILE_PATH>" {
		fmt.Println("[WARNING] Both -ip and -file should be used with -upload-cmds to generate ready-to-use upload commands.")
	}

	handlers.GenerateUploadCommands(UploadEndpoint, cfg.AttackingIP, AttackingPort, cfg.FilePath)
}

func createApplicationDirs(cfg config.ServerDirsConfig) {
	dirs := reflect.ValueOf(cfg)

	for index := 0; index < dirs.NumField(); index++ {
		if err := os.MkdirAll(dirs.Field(index).String(), 0755); err != nil {
			log.Fatalf("Failed to initialize application directories: %v", err)
		}
	}
}
