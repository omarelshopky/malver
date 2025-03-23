package config

import (
	"flag"
	"strings"
)

type Config struct {
	Port            	string
	UploadDir       	string
	DownloadDir     	string
	PingEndpoint		string
	DownloadEndpoint 	string
	UploadEndpoint   	string
	B64DecodeEndpoint 	string
	PrintUploadCommands	bool
	AttackingIP      	string
	FilePath         	string
}

func LoadConfig() Config {
	var cfg Config

	flag.StringVar(&cfg.Port, "port", "3000", "Port to run the HTTP server on")
	flag.StringVar(&cfg.UploadDir, "upload", "./uploads", "Directory for file uploads")
	flag.StringVar(&cfg.DownloadDir, "download", "./downloads", "Directory for file downloads")
	flag.StringVar(&cfg.PingEndpoint, "ping-endpoint", "/", "Endpoint for ping")
	flag.StringVar(&cfg.DownloadEndpoint, "down-endpoint", "/down", "Endpoint for file downloads")
	flag.StringVar(&cfg.UploadEndpoint, "up-endpoint", "/up", "Endpoint for file uploads")
	flag.StringVar(&cfg.B64DecodeEndpoint, "b64d-endpoint", "/b64d", "Endpoint for base64 decoding")
	flag.BoolVar(&cfg.PrintUploadCommands, "upload-cmds", false, "Generate ready-to-use upload commands")
	flag.StringVar(&cfg.AttackingIP, "ip", "<ATTACKING_IP>", "Specify the attacker's IP address (used with -upload-commands)")
	flag.StringVar(&cfg.FilePath, "file", "<FILE_PATH>", "Specify the full path of the file to be uploaded (used with -upload-commands)")

	flag.Parse()

	// Ensure DownloadEndpoint always ends with "/"
	if !strings.HasSuffix(cfg.DownloadEndpoint, "/") {
		cfg.DownloadEndpoint += "/"
	}

	return cfg
}
