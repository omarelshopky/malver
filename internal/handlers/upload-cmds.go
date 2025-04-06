package handlers

import (
	"fmt"
	"strings"

	"github.com/omarelshopky/malver/config"
)

func GenerateUploadCommands(uploadEndpoint string, attackingIP string, attackingPort string, filePath string) {
	replacer := strings.NewReplacer(
		"<ATTACKING_IP>", attackingIP,
		"<ATTACKING_PORT>", attackingPort,
		"<FILE_PATH>", filePath,
		"<UPLOAD_ENDPOINT>", uploadEndpoint,
	)

	for name, template := range config.UploadCommands {
		fmt.Printf("\n> %s\n", name)
		fmt.Println(replacer.Replace(template))
	}
}
