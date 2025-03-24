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

var logConfig *config.LoggingConfig
var logger = log.New(log.Writer(), "", 0)

type TableFormatter struct {
	maxColWidth int
}

func InitLogger(loggingConfig *config.LoggingConfig) {
	logConfig = loggingConfig
}

func LogRequest(r *http.Request, status int, message string) {
	postData := getRequestPostData(r)

	log.Printf("%s \"%s %s %s\" %d %s | %s",
		r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto,
		status, postData, message)

	tf := TableFormatter{maxColWidth: 60}

	if logConfig.Headers {
		logger.Println(tf.FormatTable("Header", "Value", r.Header))
	}

	if logConfig.Params {
		logger.Println(tf.FormatTable("Parameter", "Value", r.URL.Query()))
	}
}

func getRequestPostData(r *http.Request) string {
	postData := "-"

	if r.Method == http.MethodPost && r.Body != nil {
		bodyBytes, err := io.ReadAll(r.Body)

		if err == nil {
			postData = string(bodyBytes)
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	return postData
}

func (tf *TableFormatter) FormatTable(col1Header, col2Header string, data map[string][]string) string {
	var builder strings.Builder

	col1Width := len(col1Header)
	col2Width := len(col2Header)

	// Calculate column widths
	for key, values := range data {
		col1Width = max(col1Width, len(key))
		for _, val := range values {
			col2Width = max(col2Width, len(val))
		}
	}

	// Apply width constraints
	col1Width = min(col1Width+2, tf.maxColWidth)
	col2Width = min(col2Width+2, tf.maxColWidth)

	// Table construction components
	border := fmt.Sprintf("┌%s┬%s┐\n",
		strings.Repeat("─", col1Width),
		strings.Repeat("─", col2Width))
	header := fmt.Sprintf("│ %-*s │ %-*s │\n",
		col1Width-2, col1Header,
		col2Width-2, col2Header)
	divider := fmt.Sprintf("├%s┼%s┤\n",
		strings.Repeat("─", col1Width),
		strings.Repeat("─", col2Width))
	footer := fmt.Sprintf("└%s┴%s┘",
		strings.Repeat("─", col1Width),
		strings.Repeat("─", col2Width))

	builder.WriteString(border)
	builder.WriteString(header)
	builder.WriteString(divider)

	// Process each data entry
	for key, values := range data {
		wrappedKey := tf.wrapText(key, col1Width-2)

		for _, value := range values {
			wrappedValue := tf.wrapText(value, col2Width-2)
			maxLines := max(len(wrappedKey), len(wrappedValue))

			for i := 0; i < maxLines; i++ {
				keyPart := tf.getLine(wrappedKey, i)
				valPart := tf.getLine(wrappedValue, i)

				builder.WriteString(fmt.Sprintf("│ %-*s │ %-*s │\n",
					col1Width-2, keyPart,
					col2Width-2, valPart))
			}
		}
	}

	builder.WriteString(footer)
	return builder.String()
}

func (tf *TableFormatter) wrapText(text string, width int) []string {
	var lines []string
	remaining := text

	for len(remaining) > 0 {
		split := min(width, len(remaining))
		if spaceIdx := strings.LastIndex(remaining[:split], " "); spaceIdx > 0 {
			split = spaceIdx
		}
		lines = append(lines, remaining[:split])
		remaining = remaining[split:]
		if len(remaining) > 0 && remaining[0] == ' ' {
			remaining = remaining[1:]
		}
	}
	return lines
}

func (tf *TableFormatter) getLine(lines []string, index int) string {
	if index < len(lines) {
		return lines[index]
	}
	return ""
}
