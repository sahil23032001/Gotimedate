package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// handler writes the current date/time and hostname to the response.
func handler(w http.ResponseWriter, r *http.Request) {
	// Format the current time
	currentTime := time.Now().Format("2006-01-02 15:04:05 MST")
	
	// Get the hostname (which will be the Pod name in Kubernetes)
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	
	// Write the response, slightly adjusted for HTML formatting for readability
	fmt.Fprintf(w, "<h1>Container Status</h1><p>Current Time: <strong>%s</strong></p><p>Served by Pod: <strong>%s</strong></p>", currentTime, hostname)
}

func main() {
	// Read the port from the environment variable (standard container practice).
	// Default to 8080 if the PORT environment variable is not set.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Register the handler function for the root path
	http.HandleFunc("/", handler)
	
	log.Printf("Server starting and listening on port %s...\n", port)
	
	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
