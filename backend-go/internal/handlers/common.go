package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// WriteError writes an error response
func WriteError(w http.ResponseWriter, message string, statusCode int) {
	slog.Warn("returning error response",
		"status", statusCode,
		"message", message,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}); err != nil {
		slog.Error("failed to encode error response", "error", err)
	}
}

// NotImplemented returns a handler that responds with 501 Not Implemented
func NotImplemented(endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		WriteError(w, endpoint+" is not yet implemented in Go backend. Please use Java backend.", http.StatusNotImplemented)
	}
}
