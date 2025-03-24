package handlers

import (
	"fmt"
	"strings"
	"os"

	"gopkg.in/yaml.v3"
)

type UploadCommands struct {
	Commands map[string]string `yaml:"commands"`
}

func GenerateUploadCommands(uploadEndpoint string, attackingIP string, attackingPort string, filePath string) {
	commands, err := loadUploadCommands("config/upload-commands.yml")

	if err != nil {
		fmt.Println("Error loading upload commands:", err)

		return
	}

	replacer := strings.NewReplacer(
		"<ATTACKING_IP>", attackingIP,
		"<ATTACKING_PORT>", attackingPort,
		"<FILE_PATH>", filePath,
		"<UPLOAD_ENDPOINT>", uploadEndpoint,
	)

	for name, template := range commands.Commands {
		fmt.Printf("\n> %s\n", name)
		fmt.Println(replacer.Replace(template))
	}
}

func loadUploadCommands(filename string) (UploadCommands, error) {
	var commands UploadCommands

	data, err := os.ReadFile(filename)
	if err != nil {
		return commands, err
	}

	err = yaml.Unmarshal(data, &commands)
	if err != nil {
		return commands, err
	}

	return commands, nil
}