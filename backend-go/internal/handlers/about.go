package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dremio/dqd/internal/models"
)

// GetAbout returns version information about DQD
func GetAbout(w http.ResponseWriter, r *http.Request) {
	response := models.AboutResponse{
		Version: "0.12.3",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
