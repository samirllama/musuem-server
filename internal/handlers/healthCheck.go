package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Set JSON content type
	w.Header().Set("Content-Type", "application/json")

	// Create response
	res := map[string]string{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	}

	// Write response
	json.NewEncoder(w).Encode(res)
}
