package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// AuthRequest represents the authentication request payload
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginUser authenticates a user against the Hasura auth service
func LoginUser(data LoginData) string {
	// Build request using proper JSON marshaling
	authReq := AuthRequest{
		Username: data.Username,
		Password: data.Password,
	}

	jsonData, err := json.Marshal(authReq)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal login request")
		return ""
	}

	url := "https://auth.archon40.hasura-app.io/login"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create login request")
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send login request")
		return ""
	}
	defer resp.Body.Close()

	var r HasuraLoginData
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Error().Err(err).Msg("Failed to decode login response")
		return ""
	}

	if r.AuthToken == "" {
		log.Warn().Str("username", data.Username).Msg("Login failed - no auth token received")
		return ""
	}

	log.Info().Str("username", data.Username).Msg("User logged in successfully")
	return r.AuthToken
}
