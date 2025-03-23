package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/omarelshopky/malver/internal/logger"
)

func DownloadHandler(downloadDir string, endpointPrefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the file path after the endpoint prefix (e.g., "/down/")
		file := strings.TrimPrefix(r.URL.Path, endpointPrefix)

		if file == "" || strings.Contains(file, "..") {
			http.Error(w, "fail", http.StatusBadRequest)

			logger.LogRequest(r, http.StatusBadRequest, "invalid file path")

			return
		}

		// Create full file path and ensure it stays within downloadDir
		filePath := filepath.Join(downloadDir, filepath.Clean(file))

		if !strings.HasPrefix(filePath, filepath.Clean(downloadDir)) {
			http.Error(w, "fail", http.StatusForbidden)

			logger.LogRequest(r, http.StatusForbidden, "directory traversal attempt")

			return
		}

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.Error(w, "Not Found", http.StatusNotFound)

			logger.LogRequest(r, http.StatusNotFound, "file not found: " + file)

			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
		w.Header().Set("Content-Type", "application/octet-stream")

		http.ServeFile(w, r, filePath)

		logger.LogRequest(r, http.StatusOK, "file downloaded: " + file)
	}
}
