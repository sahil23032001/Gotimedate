package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format("Mon Jan 2 15:04:05 MST 2006")
    fmt.Fprintf(w, "<h1>Date & Time</h1><p>%s</p>", currentTime)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Running on :8080")
    http.ListenAndServe(":8080", nil)
}
