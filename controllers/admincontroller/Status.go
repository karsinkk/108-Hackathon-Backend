package admincontroller

import (
	"encoding/json"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// Status handles updating emergency status and dispatching notifications
func Status(res http.ResponseWriter, req *http.Request) {
	DB := dif.GetDB()

	var s helpers.StatusData
	if err := json.NewDecoder(req.Body).Decode(&s); err != nil {
		log.Error().Err(err).Msg("Failed to decode status request")
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	log.Debug().
		Int("emergency_id", s.Emergency_Id).
		Bool("dispatched", s.Dispatched).
		Msg("Updating emergency status")

	// Update emergency using parameterized query
	_, err := DB.Exec(
		`UPDATE emergency SET dismissed = $1, updated_description = $2 WHERE id = $3`,
		s.Dispatched, s.Updated_Description, s.Emergency_Id,
	)
	if err != nil {
		log.Error().Err(err).Int("emergency_id", s.Emergency_Id).Msg("Failed to update emergency status")
		http.Error(res, "Failed to update status", http.StatusInternalServerError)
		return
	}

	// Send notifications if dispatched
	if s.Dispatched {
		go func() {
			helpers.NotifyVehicle(s.Emergency_Id, s.User_Id)
			helpers.NotifyUser(s.Emergency_Id)
		}()
		log.Info().Int("emergency_id", s.Emergency_Id).Msg("Emergency dispatched, sending notifications")
	}

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"success": true,
		"message": "Status updated successfully",
	})
}
