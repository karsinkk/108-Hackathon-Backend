package admincontroller

import (
	"encoding/json"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// DisplayAmbulance returns all ambulance vehicles
func DisplayAmbulance(res http.ResponseWriter, req *http.Request) {
	DB := dif.GetDB()

	rows, err := DB.Query(`SELECT * FROM vehicle_data WHERE type = 1`)
	if err != nil {
		log.Error().Err(err).Msg("Failed to query ambulances")
		http.Error(res, "Failed to fetch ambulances", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	vehicles := make([]helpers.Vehicle, 0)
	for rows.Next() {
		var vehicle helpers.Vehicle
		if err := rows.Scan(
			&vehicle.Id, &vehicle.District, &vehicle.Name, &vehicle.Phone,
			&vehicle.Lat, &vehicle.Long, &vehicle.Driver, &vehicle.Vehicle_no,
			&vehicle.Username, &vehicle.Status, &vehicle.Type,
		); err != nil {
			log.Error().Err(err).Msg("Failed to scan vehicle row")
			continue
		}
		vehicles = append(vehicles, vehicle)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating vehicle rows")
	}

	log.Debug().Int("count", len(vehicles)).Msg("Retrieved ambulances")

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(vehicles)
}
