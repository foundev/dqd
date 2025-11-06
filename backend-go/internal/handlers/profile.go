package handlers

import (
	"log"
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
		log.Printf("Error parsing multipart form: %v", err)
		WriteError(w, "Failed to parse upload", http.StatusBadRequest)
		return
	}

	// Get the uploaded file
	file, fileHeader, err := r.FormFile("profile1")
	if err != nil {
		log.Printf("Error getting file from form: %v", err)
		WriteError(w, "No file uploaded with name 'profile1'", http.StatusBadRequest)
		return
	}
	defer file.Close()

	log.Printf("Processing file: %s (%d bytes)", fileHeader.Filename, fileHeader.Size)

	// Process the uploaded file (extract if archive, read if JSON)
	profileData, filename, err := fileutil.ProcessUploadedFile(fileHeader)
	if err != nil {
		log.Printf("Error processing uploaded file: %v", err)
		WriteError(w, "Failed to process uploaded file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Extracted profile from: %s (%d bytes)", filename, len(profileData))

	// Parse the profile JSON
	profileJSON, err := profile.ParseProfile(profileData)
	if err != nil {
		log.Printf("Error parsing profile JSON: %v", err)
		WriteError(w, "Failed to parse profile JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate profile
	if err := profile.ValidateProfile(profileJSON); err != nil {
		log.Printf("Invalid profile: %v", err)
		WriteError(w, "Invalid profile: "+err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Successfully parsed profile for query ID: %v", profileJSON.ID)

	// Generate simple HTML report
	htmlReport, err := profile.GenerateSimpleReport(profileJSON)
	if err != nil {
		log.Printf("Error generating report: %v", err)
		WriteError(w, "Failed to generate report: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlReport))

	duration := time.Since(start)
	log.Printf("Simple profile analysis completed in %v", duration)
}

// PostProfile handles detailed profile analysis uploads (placeholder for now)
func PostProfile(w http.ResponseWriter, r *http.Request) {
	WriteError(w, "POST /api/profile not yet implemented. Use /api/simple-profile for now.", http.StatusNotImplemented)
}

// PostProfiles handles profile comparison (placeholder for now)
func PostProfiles(w http.ResponseWriter, r *http.Request) {
	WriteError(w, "POST /api/profiles (comparison) not yet implemented", http.StatusNotImplemented)
}
