package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/dremio/dqd/internal/fileutil"
	"github.com/dremio/dqd/internal/services/profile"
)

const maxUploadSize = 10 * 1024 * 1024 // 10MB

// PostSimpleProfile handles simple profile analysis uploads
func PostSimpleProfile(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Parse multipart form
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		slog.Error("failed to parse multipart form", "error", err)
		WriteError(w, "Failed to parse upload", http.StatusBadRequest)
		return
	}

	// Get the uploaded file
	file, fileHeader, err := r.FormFile("profile1")
	if err != nil {
		slog.Error("failed to get file from form", "error", err)
		WriteError(w, "No file uploaded with name 'profile1'", http.StatusBadRequest)
		return
	}

	// Process file and ensure proper cleanup
	profileData, filename, err := func() ([]byte, string, error) {
		defer func() {
			if closeErr := file.Close(); closeErr != nil {
				slog.Error("failed to close uploaded file",
					"filename", fileHeader.Filename,
					"error", closeErr,
				)
			}
		}()

		slog.Info("processing uploaded file",
			"filename", fileHeader.Filename,
			"size", fileHeader.Size,
		)

		// Process the uploaded file (extract if archive, read if JSON)
		data, name, err := fileutil.ProcessUploadedFile(fileHeader)
		if err != nil {
			return nil, "", fmt.Errorf("processing uploaded file: %w", err)
		}

		return data, name, nil
	}()

	if err != nil {
		slog.Error("failed to process upload", "error", err, "filename", fileHeader.Filename)
		WriteError(w, "Failed to process uploaded file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	slog.Info("extracted profile",
		"original_filename", fileHeader.Filename,
		"extracted_filename", filename,
		"size", len(profileData),
	)

	// Parse the profile JSON
	profileJSON, err := profile.ParseProfile(profileData)
	if err != nil {
		slog.Error("failed to parse profile JSON", "error", err, "filename", filename)
		WriteError(w, "Failed to parse profile JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate profile
	if err := profile.ValidateProfile(profileJSON); err != nil {
		slog.Error("invalid profile", "error", err, "filename", filename)
		WriteError(w, "Invalid profile: "+err.Error(), http.StatusBadRequest)
		return
	}

	slog.Info("successfully parsed profile",
		"query_id", profileJSON.ID,
		"user", profileJSON.User,
	)

	// Generate simple HTML report
	htmlReport, err := profile.GenerateSimpleReport(profileJSON)
	if err != nil {
		slog.Error("failed to generate report", "error", err)
		WriteError(w, "Failed to generate report: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(htmlReport)); err != nil {
		slog.Error("failed to write response", "error", err)
		return
	}

	duration := time.Since(start)
	slog.Info("request completed",
		"method", r.Method,
		"path", r.URL.Path,
		"duration", duration,
		"filename", fileHeader.Filename,
	)
}

// PostProfile handles detailed profile analysis uploads (placeholder for now)
func PostProfile(w http.ResponseWriter, r *http.Request) {
	WriteError(w, "POST /api/profile not yet implemented. Use /api/simple-profile for now.", http.StatusNotImplemented)
}

// PostProfiles handles profile comparison (placeholder for now)
func PostProfiles(w http.ResponseWriter, r *http.Request) {
	WriteError(w, "POST /api/profiles (comparison) not yet implemented", http.StatusNotImplemented)
}
