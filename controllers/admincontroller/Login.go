package admincontroller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// Login handles admin user authentication
func Login(res http.ResponseWriter, req *http.Request) {
	var r helpers.LoginData
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		log.Error().Err(err).Msg("Failed to decode login request")
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	log.Debug().Str("username", r.Username).Msg("Admin login attempt")

	var data helpers.LoginReturnData
	data.Auth = helpers.LoginUser(r)

	if data.Auth == "" {
		log.Warn().Str("username", r.Username).Msg("Login failed - invalid credentials")
		http.Error(res, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Extract admin user ID from username
	var adminID int64
	if strings.Contains(r.Username, "adminuser") {
		parts := strings.Split(r.Username, "adminuser")
		if len(parts) > 1 {
			adminID, _ = strconv.ParseInt(parts[1], 10, 64)
		}
	}

	log.Info().Str("username", r.Username).Int64("admin_id", adminID).Msg("Admin logged in successfully")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"auth":     data.Auth,
		"admin_id": adminID,
	})
}
