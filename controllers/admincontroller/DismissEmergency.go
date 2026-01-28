package admincontroller

import (
	"encoding/json"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// DismissEmergency handles dismissing an emergency
func DismissEmergency(res http.ResponseWriter, req *http.Request) {
	DB := dif.GetDB()

	var r helpers.DismissEmergencyData
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		log.Error().Err(err).Msg("Failed to decode dismiss emergency request")
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	log.Debug().
		Int("emergency_id", r.Emergency_Id).
		Str("reason", r.Dismissed_Reason).
		Msg("Dismissing emergency")

	// Update emergency status using parameterized query
	_, err := DB.Exec(
		`UPDATE emergency
		 SET status = false, dismissed = true, updated_description = $1
		 WHERE id = $2`,
		r.Dismissed_Reason, r.Emergency_Id,
	)
	if err != nil {
		log.Error().Err(err).Int("emergency_id", r.Emergency_Id).Msg("Failed to update emergency status")
		http.Error(res, "Failed to dismiss emergency", http.StatusInternalServerError)
		return
	}

	// Delete dispatched vehicles for this emergency
	_, err = DB.Exec(
		`DELETE FROM dispatched_vehicles WHERE emergency_id = $1`,
		r.Emergency_Id,
	)
	if err != nil {
		log.Error().Err(err).Int("emergency_id", r.Emergency_Id).Msg("Failed to delete dispatched vehicles")
	}

	log.Info().Int("emergency_id", r.Emergency_Id).Msg("Emergency dismissed")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"success": true,
		"message": "Emergency dismissed successfully",
	})
}
