package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// Emergency handles emergency request submissions from users
func Emergency(res http.ResponseWriter, req *http.Request) {
	var u helpers.EmergencyUserData
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		log.Error().Err(err).Msg("Failed to decode emergency request")
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	log.Info().
		Str("lat", u.Lat).
		Str("long", u.Long).
		Str("name", u.Name).
		Int("type", u.Type).
		Msg("New emergency request received")

	vehiclesData := make([]helpers.VehicleData, 0)

	// For ambulance emergencies (Type 1), find closest vehicles
	if u.Type == 1 {
		vehiclesData = helpers.GetClosest(u.Lat, u.Long, u.Type, u.Number)
		if vehiclesData == nil {
			log.Warn().Msg("No vehicles available for emergency")
		}
	}

	// Add emergency to database
	if err := helpers.AddEmergency(vehiclesData, u); err != nil {
		log.Error().Err(err).Msg("Failed to add emergency")
		http.Error(res, "Failed to process emergency", http.StatusInternalServerError)
		return
	}

	log.Info().
		Int("vehicles_dispatched", len(vehiclesData)).
		Int("emergency_type", u.Type).
		Msg("Emergency processed successfully")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"success":            true,
		"message":            "Your Emergency Vehicle is on its way.",
		"vehicles_dispatched": len(vehiclesData),
	})
}
