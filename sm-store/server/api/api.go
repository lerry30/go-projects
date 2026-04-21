package api

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Error string
}

// JSON response function
func ResponseJSON(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}
