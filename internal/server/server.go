package server

import (
	"log"
	"net/http"

	"github.com/omarelshopky/malver/config"
	"github.com/omarelshopky/malver/internal/handlers"
)

func Start(port string, endpoints config.EndpointsConfig, dirs config.ServerDirsConfig) {
	http.HandleFunc(endpoints.Ping, handlers.PingHandler)
	http.HandleFunc(endpoints.Download, handlers.DownloadHandler(dirs.Download, endpoints.Download))
	http.HandleFunc(endpoints.Upload, handlers.UploadHandler(dirs.Upload))
	http.HandleFunc(endpoints.B64Decode, handlers.B64DecodeHandler)

	log.Printf("Starting server on :%s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
