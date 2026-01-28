package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// RegisterUser creates a new user account in the Hasura auth service
func RegisterUser(data AdminRegisterData) string {
	// Build request using proper JSON marshaling
	// Note: Using username as password (preserved from original behavior)
	authReq := AuthRequest{
		Username: data.Username,
		Password: data.Username,
	}

	jsonData, err := json.Marshal(authReq)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal signup request")
		return ""
	}

	url := "https://auth.archon40.hasura-app.io/signup"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create signup request")
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send signup request")
		return ""
	}
	defer resp.Body.Close()

	var r HasuraSignupData
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Error().Err(err).Msg("Failed to decode signup response")
		return ""
	}

	if r.AuthToken == "" {
		log.Warn().Str("username", data.Username).Msg("Signup failed - no auth token received")
		return ""
	}

	log.Info().Str("username", data.Username).Msg("User registered successfully")
	return r.AuthToken
}
