package admincontroller

import (
	"encoding/json"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

// ModifySeen marks emergencies as seen
func ModifySeen(res http.ResponseWriter, req *http.Request) {
	DB := dif.GetDB()

	var u helpers.SeenData
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		log.Error().Err(err).Msg("Failed to decode modify seen request")
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if len(u.Id) == 0 {
		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("108", "An NP-Incomplete Project")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]interface{}{
			"success": true,
			"message": "No IDs to update",
		})
		return
	}

	log.Debug().Ints("emergency_ids", u.Id).Msg("Marking emergencies as seen")

	// Use parameterized query with array
	_, err := DB.Exec(
		`UPDATE emergency SET seen = true WHERE id = ANY($1)`,
		pq.Array(u.Id),
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update emergency seen status")
		http.Error(res, "Failed to update status", http.StatusInternalServerError)
		return
	}

	log.Info().Ints("emergency_ids", u.Id).Msg("Emergencies marked as seen")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"success": true,
		"updated": len(u.Id),
	})
}
