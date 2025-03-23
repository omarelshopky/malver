package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/omarelshopky/malver/internal/logger"
)

// handleMultipartUpload processes multipart/form-data uploads.
func handleMultipartUpload(r *http.Request) (io.ReadCloser, string, error) {
	file, handler, err := r.FormFile("file")

	if err != nil {
		return nil, "", fmt.Errorf("failed to get file from form: %w", err)
	}

	return file, handler.Filename, nil
}

// handleOctetStreamUpload processes application/octet-stream uploads.
func handleOctetStreamUpload(r *http.Request) (io.ReadCloser, string) {
	filename := r.Header.Get("Filename")

	if filename == "" {
		filename = fmt.Sprintf("upload_%d", time.Now().UnixNano())
	}

	return r.Body, filename
}

func saveFile(uploadDir string, file io.Reader, filename string) (string, error) {
	safeFilename := filepath.Base(filename)
	filePath := filepath.Join(uploadDir, safeFilename)

	absUploadDir, err := filepath.Abs(uploadDir)
	if err != nil {
		return "", err
	}

	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	// Ensure the final file path is inside the upload directory
	if !strings.HasPrefix(absFilePath, absUploadDir) {
		return "", os.ErrPermission
	}

	out, err := os.Create(absFilePath)

	if err != nil {
		return "", err
	}

	defer out.Close()

	_, err = io.Copy(out, file)

	if err != nil {
		return "", err
	}

	return absFilePath, nil
}

func UploadHandler(uploadDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var file io.ReadCloser
		var filename string
		var err error

		contentType := r.Header.Get("Content-Type")

		if strings.Contains(contentType, "multipart/form-data") {
			file, filename, err = handleMultipartUpload(r)
		} else {
			file, filename = handleOctetStreamUpload(r)
		}

		if err != nil {
			http.Error(w, "fail", http.StatusBadRequest)

			logger.LogRequest(r, http.StatusBadRequest, err.Error())

			return
		}

		defer file.Close()

		savedPath, err := saveFile(uploadDir, file, filename)

		if err != nil {
			http.Error(w, "fail", http.StatusInternalServerError)

			logger.LogRequest(r, http.StatusInternalServerError, "failed to save file")

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))

		logger.LogRequest(r, http.StatusOK, "file uploaded: " + filepath.Base(savedPath))
	}
}
