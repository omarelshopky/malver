package config

import (
	"flag"
	"strings"
)

type Config struct {
	Server              ServerConfig
	Endpoints           EndpointsConfig
	Logging             LoggingConfig
	PrintUploadCommands PrintUploadCommandsConfig
}

type ServerConfig struct {
	Port string
	Dirs ServerDirsConfig
}

type ServerDirsConfig struct {
	Upload   string
	Download string
}

type EndpointsConfig struct {
	Ping      string
	Download  string
	Upload    string
	B64Decode string
}

type LoggingConfig struct {
	Headers bool
	Params  bool
}

type PrintUploadCommandsConfig struct {
	Enabled     bool
	AttackingIP string
	FilePath    string
}

func LoadConfig() Config {
	var cfg Config

	setupServerFlags(&cfg.Server)
	setupEndpointFlags(&cfg.Endpoints)
	setupLoggingFlags(&cfg.Logging)
	setupUploadCommandFlags(&cfg.PrintUploadCommands)

	flag.Parse()

	normalizeEndpoints(&cfg.Endpoints)

	return cfg
}

func setupServerFlags(cfg *ServerConfig) {
	flag.StringVar(&cfg.Port, "port", "3000", "Port to run the HTTP server on")
	flag.StringVar(&cfg.Dirs.Upload, "upload", "./uploads", "Directory for file uploads")
	flag.StringVar(&cfg.Dirs.Download, "download", "./downloads", "Directory for file downloads")
}

func setupEndpointFlags(cfg *EndpointsConfig) {
	flag.StringVar(&cfg.Ping, "ping-endpoint", "/", "Endpoint for ping")
	flag.StringVar(&cfg.Download, "down-endpoint", "/down", "Endpoint for file downloads")
	flag.StringVar(&cfg.Upload, "up-endpoint", "/up", "Endpoint for file uploads")
	flag.StringVar(&cfg.B64Decode, "b64d-endpoint", "/b64d", "Endpoint for base64 decoding")
}

func setupLoggingFlags(cfg *LoggingConfig) {
	flag.BoolVar(&cfg.Headers, "headers", false, "Log request headers")
	flag.BoolVar(&cfg.Params, "params", false, "Log request query parameters")
}

func setupUploadCommandFlags(cfg *PrintUploadCommandsConfig) {
	flag.BoolVar(&cfg.Enabled, "upload-cmds", false, "Generate ready-to-use upload commands")
	flag.StringVar(&cfg.AttackingIP, "ip", "<ATTACKING_IP>", "Attacker's IP address (used with -upload-cmds)")
	flag.StringVar(&cfg.FilePath, "file", "<FILE_PATH>", "Full path of file to upload (used with -upload-cmds)")
}

func normalizeEndpoints(cfg *EndpointsConfig) {
	endpoints := map[*string]bool{
		&cfg.Ping:      false,
		&cfg.Download:  true,
		&cfg.Upload:    false,
		&cfg.B64Decode: false,
	}

	for endpoint, needsTrailingSlash := range endpoints {
		ensureEndpointFormat(endpoint, needsTrailingSlash)
	}
}

func ensureEndpointFormat(endpoint *string, trailingSlash bool) {
	trimmedEndpoint := strings.Trim(strings.TrimSpace(*endpoint), "/")

	if !trailingSlash {
		*endpoint = "/" + trimmedEndpoint
		return
	}

	*endpoint = "/" + trimmedEndpoint + "/"
}
