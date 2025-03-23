package server

import (
	"log"
	"net/http"

	"github.com/omarelshopky/malver/internal/handlers"
	"github.com/omarelshopky/malver/config"
)

func Start(cfg config.Config) {
	http.HandleFunc(cfg.PingEndpoint, handlers.PingHandler)
	http.HandleFunc(cfg.DownloadEndpoint, handlers.DownloadHandler(cfg.DownloadDir, cfg.DownloadEndpoint))
	http.HandleFunc(cfg.UploadEndpoint, handlers.UploadHandler(cfg.UploadDir))
	http.HandleFunc(cfg.B64DecodeEndpoint, handlers.B64DecodeHandler)

	log.Printf("Starting server on :%s", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
