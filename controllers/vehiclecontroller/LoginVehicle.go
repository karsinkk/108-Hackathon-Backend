package vehiclecontroller

import (
	"encoding/json"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// LoginVehicle handles vehicle authentication and location update
func LoginVehicle(res http.ResponseWriter, req *http.Request) {
	DB := dif.GetDB()

	var a helpers.VehicleLoginData
	if err := json.NewDecoder(req.Body).Decode(&a); err != nil {
		log.Error().Err(err).Msg("Failed to decode vehicle login request")
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	log.Debug().Str("username", a.Username).Msg("Vehicle login attempt")

	// Authenticate vehicle
	loginData := helpers.LoginData{Username: a.Username, Password: a.Password}
	auth := helpers.LoginUser(loginData)

	var v helpers.Vehicle_Id
	if auth == "" {
		log.Warn().Str("username", a.Username).Msg("Vehicle login failed - invalid credentials")
		http.Error(res, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Get vehicle ID using parameterized query
	err := DB.QueryRow(
		`SELECT id FROM vehicle_data WHERE username = $1`,
		a.Username,
	).Scan(&v.Id)
	if err != nil {
		log.Error().Err(err).Str("username", a.Username).Msg("Failed to get vehicle ID")
		http.Error(res, "Vehicle not found", http.StatusNotFound)
		return
	}

	// Update vehicle location using parameterized query
	_, err = DB.Exec(
		`UPDATE vehicle_data SET lat = $1, long = $2, time = NOW()::timestamp WHERE id = $3`,
		a.Lat, a.Long, v.Id,
	)
	if err != nil {
		log.Error().Err(err).Int("vehicle_id", v.Id).Msg("Failed to update vehicle location")
	}

	// Update vehicle token using parameterized query
	_, err = DB.Exec(
		`UPDATE vehicle_token_data SET token = $1 WHERE vehicle_id = $2`,
		a.Token, v.Id,
	)
	if err != nil {
		log.Error().Err(err).Int("vehicle_id", v.Id).Msg("Failed to update vehicle token")
	}

	log.Info().Str("username", a.Username).Int("vehicle_id", v.Id).Msg("Vehicle logged in successfully")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(v)
}
