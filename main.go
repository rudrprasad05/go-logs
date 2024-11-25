package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rudrprasad05/go-logs/logs"
)

func main() {
	// Create a new logger instance
	logger, err := logs.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	// Ensure the logger is closed when the app exits
	defer logger.Close()

	// Define your HTTP routes and use the logging middleware
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request to /")
		fmt.Fprintln(w, "Hello, World!")
	})

	// Wrap the HTTP router with the logging middleware
	loggedHandler := logs.LoggingMiddleware(logger, router)

	// Start the HTTP server
	http.ListenAndServe(":8080", loggedHandler)
}
