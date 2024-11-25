# go-logs

A lightweight and developer-friendly logging utility for Go projects. The `go-logs` package simplifies logging HTTP requests, debug messages, errors, and more while ensuring clear, structured logs for your applications.

---

## **Why I Built This**

`go-logs` was built to streamline logging in my Go projects, particularly for scenarios where I needed:

- A consistent, structured log format.
- Easy-to-use middleware for HTTP request logging.
- Automatic management of log files, including daily rotation.
- Minimal setup effort for integrating logging into my projects.

By using this logger, I can focus on building application features without worrying about maintaining logging code.

---

## **Features**

- Logs messages with different levels: `INFO`, `DEBUG`, `ERROR`, and `HTTP`.
- Automatically rotates log files daily.
- Includes an HTTP middleware for logging request details (method, path, and client IP).
- Writes logs to both console and files for easy debugging and audit purposes.
- Ensures the `logs` directory is ignored by Git automatically to avoid cluttering repositories.

---

## **Installation**

```bash
go get github.com/rudrprasad05/go-logs
```

## **Examples**

```go
package main

import (
	"github.com/rudrprasad05/go-logs"
)

func main() {
	logger, err := logs.NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Info("This is an info log.")
	logger.Debug("This is a debug log.")
	logger.Error("This is an error log.")
}
```

```go
package main

import (
	"github.com/rudrprasad05/go-logs"
	"net/http"
)

func main() {
	logger, err := logs.NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	loggedRouter := logs.LoggingMiddleware(logger, mux)
	http.ListenAndServe(":8080", loggedRouter)
}

```

## **Extra Features**

- Automatic Git Ignore: Adds /logs to .gitignore to prevent committing log files. If a .gitignore file doesn't exist, it creates one.
- Daily Rotation: Log files are automatically rotated based on the current date (e.g., logs/22-11-24.log).
- Request IP Handling: Accurately logs client IPs, accounting for X-Forwarded-For headers when present.
