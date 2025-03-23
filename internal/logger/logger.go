package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/omarelshopky/malver/config"
)


var cfg config.Config

var headerLogger = log.New(log.Writer(), "", 0)

func InitLogger(c config.Config) {
	cfg = c
}

func LogRequest(r *http.Request, status int, message string) {
	postData := "-"

	if r.Method == http.MethodPost && r.Body != nil {
		bodyBytes, err := io.ReadAll(r.Body)
		if err == nil {
			postData = string(bodyBytes)
			// Restore the request body since it gets read
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	log.Printf("%s \"%s %s %s\" %d %s | %s", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto, status, postData, message)

	if cfg.LogHeaders {
		headerLogger.Println(formatHeadersTable(r.Header))
	}
}

func formatHeadersTable(headers http.Header) string {
	var builder strings.Builder

	maxKeyLen := 0
	maxValLen := 0

	for key, values := range headers {
		maxKeyLen = max(maxKeyLen, len(key))

		for _, value := range values {
			maxValLen = max(maxValLen, len(value))
		}
	}

	maxKeyLen = min(maxKeyLen + 2, 60)
	maxValLen = min(maxValLen + 2, 60)

	border := fmt.Sprintf("┌%s┬%s┐\n", strings.Repeat("─", maxKeyLen), strings.Repeat("─", maxValLen))
	headerRow := fmt.Sprintf("│ %-*s │ %-*s │\n", maxKeyLen-2, "Header", maxValLen-2, "Value")
	divider := fmt.Sprintf("├%s┼%s┤\n", strings.Repeat("─", maxKeyLen), strings.Repeat("─", maxValLen))
	footer := fmt.Sprintf("└%s┴%s┘\n", strings.Repeat("─", maxKeyLen), strings.Repeat("─", maxValLen))

	builder.WriteString(border)
	builder.WriteString(headerRow)
	builder.WriteString(divider)

	for key, values := range headers {
		wrappedKey := wrapText(key, maxKeyLen-2)

		for _, value := range values {
			wrappedValue := wrapText(value, maxValLen-2)

			maxRows := max(len(wrappedKey), len(wrappedValue))

			for i := 0; i < maxRows; i++ {
				keyPart := ""
				if i < len(wrappedKey) {
					keyPart = wrappedKey[i]
				}

				valuePart := ""
				if i < len(wrappedValue) {
					valuePart = wrappedValue[i]
				}

				builder.WriteString(fmt.Sprintf("│ %-*s │ %-*s │\n", maxKeyLen-2, keyPart, maxValLen-2, valuePart))
			}
		}
	}

	builder.WriteString(footer)

	return builder.String()
}

func wrapText(text string, width int) []string {
	var lines []string

	for len(text) > width {
		split := width

		if space := strings.LastIndex(text[:width], " "); space > 0 {
			split = space
		}

		lines = append(lines, text[:split])
		text = text[split:]
	}

	lines = append(lines, text)

	return lines
}