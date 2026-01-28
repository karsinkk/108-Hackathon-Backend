package admincontroller

import (
	"encoding/json"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// DismissVehicle removes a vehicle from a dispatched emergency
func DismissVehicle(res http.ResponseWriter, req *http.Request) {
	DB := dif.GetDB()

	var r helpers.DismissData
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		log.Error().Err(err).Msg("Failed to decode dismiss vehicle request")
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	log.Debug().
		Int("vehicle_id", r.Vehicle_Id).
		Int("emergency_id", r.Emergency_Id).
		Msg("Dismissing vehicle from emergency")

	// Delete using parameterized query
	result, err := DB.Exec(
		`DELETE FROM dispatched_vehicles WHERE vehicle_id = $1 AND emergency_id = $2`,
		r.Vehicle_Id, r.Emergency_Id,
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to dismiss vehicle")
		http.Error(res, "Failed to dismiss vehicle", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	log.Info().
		Int("vehicle_id", r.Vehicle_Id).
		Int("emergency_id", r.Emergency_Id).
		Int64("rows_affected", rowsAffected).
		Msg("Vehicle dismissed from emergency")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"success": true,
		"message": "Vehicle dismissed successfully",
	})
}
