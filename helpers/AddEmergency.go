package helpers

import (
	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/rs/zerolog/log"
)

// AddEmergency creates a new emergency record and dispatches vehicles
func AddEmergency(vehiclesData []VehicleData, u EmergencyUserData) error {
	DB := dif.GetDB()

	// Insert emergency using parameterized query
	var emergencyID int
	err := DB.QueryRow(
		`INSERT INTO emergency(lat, long, name, phone, type, description)
		 VALUES($1, $2, $3, $4, $5, $6)
		 RETURNING id`,
		u.Lat, u.Long, u.Name, u.Phone, u.Type, u.Description,
	).Scan(&emergencyID)

	if err != nil {
		log.Error().Err(err).Msg("Failed to insert emergency")
		return err
	}

	log.Info().Int("emergency_id", emergencyID).Msg("Emergency created")

	// Dispatch vehicles for ambulance emergencies (Type 1)
	if u.Type == 1 {
		for _, v := range vehiclesData {
			_, err := DB.Exec(
				`INSERT INTO dispatched_vehicles(emergency_id, vehicle_id, time_taken, distance)
				 VALUES($1, $2, $3, $4)`,
				emergencyID, v.Id, v.Time, v.Distance,
			)
			if err != nil {
				log.Error().
					Err(err).
					Int("emergency_id", emergencyID).
					Int("vehicle_id", v.Id).
					Msg("Failed to dispatch vehicle")
			}
		}
	}

	// Store emergency token
	_, err = DB.Exec(
		`INSERT INTO emergency_token_data(emergency_id, token) VALUES($1, $2)`,
		emergencyID, u.Token,
	)
	if err != nil {
		log.Error().
			Err(err).
			Int("emergency_id", emergencyID).
			Msg("Failed to store emergency token")
		return err
	}

	return nil
}
