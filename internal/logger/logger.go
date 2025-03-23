package logger

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func LogRequest(r *http.Request, status int, message string) {
	postData := "-"

	if r.Method == http.MethodPost {
		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)

			if err == nil {
				postData = string(bodyBytes)
				// Restore the request body since it gets read
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}
	}

	log.Printf("%s \"%s %s %s\" %d %s - %s", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto, status, postData, message)
}
