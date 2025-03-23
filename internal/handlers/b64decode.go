package handlers

import (
	"encoding/base64"
	"net/http"

	"github.com/omarelshopky/malver/internal/logger"
)

func B64DecodeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("d")
	decoded, err := base64.StdEncoding.DecodeString(query)

	if err != nil {
		http.Error(w, "fail", http.StatusBadRequest)

		logger.LogRequest(r, http.StatusBadRequest, "invalid base64 string")

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))

	logger.LogRequest(r, http.StatusOK, "decoded: " + string(decoded))
}