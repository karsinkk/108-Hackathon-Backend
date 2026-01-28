package admincontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// AddVehicle handles adding a new vehicle to the system
func AddVehicle(res http.ResponseWriter, req *http.Request) {
	DB := dif.GetDB()

	var r helpers.VehicleAddData
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		log.Error().Err(err).Msg("Failed to decode add vehicle request")
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	log.Debug().
		Str("vehicle_no", r.Vehicle_no).
		Str("driver", r.Driver).
		Msg("Adding new vehicle")

	// Insert vehicle using parameterized query
	var id int
	err := DB.QueryRow(
		`INSERT INTO vehicle_data(vehicle_no, driver, phone, type)
		 VALUES($1, $2, $3, 1)
		 RETURNING id`,
		r.Vehicle_no, r.Driver, r.Phone,
	).Scan(&id)

	if err != nil {
		log.Error().Err(err).Msg("Failed to insert vehicle")
		http.Error(res, "Failed to add vehicle", http.StatusInternalServerError)
		return
	}

	// Generate username based on vehicle ID
	username := fmt.Sprintf("vehicle%d", id)

	// Update vehicle with username
	_, err = DB.Exec(
		`UPDATE vehicle_data SET username = $1 WHERE id = $2`,
		username, id,
	)
	if err != nil {
		log.Error().Err(err).Int("vehicle_id", id).Msg("Failed to update vehicle username")
	}

	// Insert token record
	_, err = DB.Exec(
		`INSERT INTO vehicle_token_data(vehicle_id, token) VALUES($1, '')`,
		id,
	)
	if err != nil {
		log.Error().Err(err).Int("vehicle_id", id).Msg("Failed to create vehicle token record")
	}

	// Register user in auth system
	data := helpers.AdminRegisterData{Username: username}
	helpers.RegisterUser(data)

	log.Info().
		Int("vehicle_id", id).
		Str("username", username).
		Msg("Vehicle added successfully")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"id":       id,
		"username": username,
	})
}
