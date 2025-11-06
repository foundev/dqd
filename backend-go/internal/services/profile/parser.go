package profile

import (
	"encoding/json"
	"fmt"

	"github.com/dremio/dqd/internal/models"
)

// ParseProfile parses a profile JSON byte array into a ProfileJSON struct
func ParseProfile(data []byte) (*models.ProfileJSON, error) {
	var profile models.ProfileJSON
	if err := json.Unmarshal(data, &profile); err != nil {
		return nil, fmt.Errorf("failed to parse profile JSON: %w", err)
	}
	return &profile, nil
}

// ValidateProfile performs basic validation on a profile
func ValidateProfile(profile *models.ProfileJSON) error {
	if profile == nil {
		return fmt.Errorf("profile is nil")
	}
	if profile.Query == "" {
		return fmt.Errorf("profile has no query")
	}
	return nil
}
