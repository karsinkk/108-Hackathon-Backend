package admincontroller

import (
	"encoding/json"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// Register handles admin user registration
func Register(res http.ResponseWriter, req *http.Request) {
	var r helpers.AdminRegisterData
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		log.Error().Err(err).Msg("Failed to decode register request")
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	log.Debug().Str("username", r.Username).Msg("Registering new admin user")

	var data helpers.SignUpReturnData
	data.Auth_token = helpers.RegisterUser(r)

	if data.Auth_token == "" {
		log.Warn().Str("username", r.Username).Msg("Registration failed")
		http.Error(res, "Registration failed", http.StatusInternalServerError)
		return
	}

	log.Info().Str("username", r.Username).Msg("Admin user registered successfully")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(data)
}
