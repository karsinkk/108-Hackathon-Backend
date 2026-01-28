package vehiclecontroller

import (
	"encoding/json"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// Finish handles completing an emergency response
func Finish(res http.ResponseWriter, req *http.Request) {
	DB := dif.GetDB()

	var a helpers.Vehicle_Id
	if err := json.NewDecoder(req.Body).Decode(&a); err != nil {
		log.Error().Err(err).Msg("Failed to decode finish request")
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	log.Debug().Int("vehicle_id", a.Id).Msg("Vehicle finishing emergency")

	// Get emergency ID for this vehicle using parameterized query
	var emergencyID int
	err := DB.QueryRow(
		`SELECT emergency_id FROM dispatched_vehicles WHERE vehicle_id = $1`,
		a.Id,
	).Scan(&emergencyID)
	if err != nil {
		log.Error().Err(err).Int("vehicle_id", a.Id).Msg("Failed to get emergency ID")
		http.Error(res, "Failed to find emergency", http.StatusNotFound)
		return
	}

	// Update emergency status
	_, err = DB.Exec(
		`UPDATE emergency SET status = false WHERE id = $1`,
		emergencyID,
	)
	if err != nil {
		log.Error().Err(err).Int("emergency_id", emergencyID).Msg("Failed to update emergency status")
	}

	// Make vehicle available again
	_, err = DB.Exec(
		`UPDATE vehicle_data SET status = true WHERE id = $1`,
		a.Id,
	)
	if err != nil {
		log.Error().Err(err).Int("vehicle_id", a.Id).Msg("Failed to update vehicle status")
	}

	log.Info().
		Int("vehicle_id", a.Id).
		Int("emergency_id", emergencyID).
		Msg("Emergency completed")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"success":      true,
		"emergency_id": emergencyID,
	})
}
