package handlers

import (
	"net/http"

	"github.com/omarelshopky/malver/internal/logger"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))

	logger.LogRequest(r, http.StatusOK, "ping")
}