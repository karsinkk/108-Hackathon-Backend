package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// SubmitRating handles rating submissions for hospitals
func SubmitRating(res http.ResponseWriter, req *http.Request) {
	DB := dif.GetDB()

	var u helpers.Rating
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		log.Error().Err(err).Msg("Failed to decode rating request")
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	log.Debug().Int("hospital_id", u.Id).Int("rating", u.Rating).Msg("Submitting rating")

	// Update hospital rating using parameterized query
	_, err := DB.Exec(
		`UPDATE hospital
		 SET rating = ((rating * rating_count + $1) / (rating_count + 1)),
		     rating_count = rating_count + 1
		 WHERE id = $2`,
		u.Rating, u.Id,
	)
	if err != nil {
		log.Error().Err(err).Int("hospital_id", u.Id).Msg("Failed to update hospital rating")
		http.Error(res, "Failed to submit rating", http.StatusInternalServerError)
		return
	}

	log.Info().Int("hospital_id", u.Id).Int("rating", u.Rating).Msg("Rating submitted successfully")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"success": true,
		"message": "Rating submitted successfully",
	})
}
