package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HellHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method                      // Get HTTP method (GET, POST, etc.)
	userAgent := r.Header.Get("User-Agent") // Get request headers
	name := r.URL.Query().Get("name")       // Get the 'name' from query parameter (?name=John)
	if name == "" {
		name = "Guest"
	}

	// Write personalized response
	w.Write([]byte("Welcome to hell!" + name + "😈"))

	fmt.Printf("Hellhandler logging http Request method %v \n", method)
	fmt.Printf("Hellhandler logging http Request userAgent %v \n", userAgent)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Log incoming request
	log.Printf("Received request to %s from %s", r.URL.Path, r.RemoteAddr)

	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method alloewd", http.StatusMethodNotAllowed)
		return
	}

	// Use of Query Param to get values from the URL
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	response := map[string]string{
		"message": "Hello, " + name + "👋",
		"status":  "success",
	}

	w.Header().Set("Content-Type", "application/json") // Set content type header

	json.NewEncoder(w).Encode(response) // Write response

}
