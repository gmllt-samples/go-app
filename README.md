# go-app – Simple JSON HTTP test server

This Go app is a simple HTTP server that sends JSON responses.  
You can use it to test timeouts, HTTP status codes, and response sizes.  
Great for testing proxies, load balancers, or HTTP clients.

---

## Quick Start

### Requirements

- Go 1.21 or newer
- `make` installed

### Build and run

```bash
make run
```

The server runs at:

    http://localhost:8080

---

## Parameters

You can pass query parameters to control the response:

| Parameter       | Description                      | Examples                            |
|-----------------|----------------------------------|-------------------------------------|
| `wait`          | Wait time before responding      | `1s`, `500ms`, `1s-2s`, `1s,2s`     |
| `status`        | HTTP status code                 | `200`, `404`, `200-299`, `200,404`  |
| `response_size` | Size of response body (in bytes) | `100`, `1K`, `10K-20K`, `1K,5K,10K` |

All parameters support:

- **single value**
- **range** using `-` (e.g. `100-200`)
- **list** using `,` (e.g. `200,404,500`)

---

## Example `curl` requests

### Wait 2 seconds before responding:

```bash
curl "http://localhost:8080?wait=2s"
```

### Random HTTP status code between 200 and 299:

```bash
curl "http://localhost:8080?status=200-299"
```

### Fixed 404 with 5KB body:

```bash
curl "http://localhost:8080?status=404&response_size=5K"
```

### Random response size from list:

```bash
curl "http://localhost:8080?response_size=1K,2K,5K"
```

### Full example: slow 500 error with 10KB body

```bash
curl "http://localhost:8080?wait=1s&status=500&response_size=10K"
```

---

## Makefile usage

### Build the app:

```bash
make build
```

Output binary goes to: `./bin/server`

### Run the app (on port 8080):

```bash
make run
```

You can also override the port:

```bash
PORT=3000 make run
```

### Clean build files:

```bash
make clean
```

---

## Project Structure

```
go-app/
├── cmd/
│   └── server/          # Main app entry (main.go)
├── internal/
│   ├── app/             # App logic and handler
│   ├── log/             # JSON logger
│   ├── parser/          # wait, size, status parsing
│   └── response/        # JSON response struct
├── bin/                 # Compiled binaries (via make build)
├── go.mod
├── Makefile
└── README.md
```

