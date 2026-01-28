package helpers

import (
	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/rs/zerolog/log"
)

// NotifyVehicle updates dispatched vehicles and sends notifications
func NotifyVehicle(emergencyID int, userID int) {
	DB := dif.GetDB()

	// Update dispatched vehicles with user ID
	_, err := DB.Exec(
		`UPDATE dispatched_vehicles SET user_id = $1 WHERE emergency_id = $2`,
		userID, emergencyID,
	)
	if err != nil {
		log.Error().Err(err).Int("emergency_id", emergencyID).Msg("Failed to update dispatched vehicles")
	}

	// Get emergency details
	vehicleNotification := VehicleNotificationData{}
	err = DB.QueryRow(
		`SELECT lat, long, name, phone, type, updated_description
		 FROM emergency WHERE id = $1`,
		emergencyID,
	).Scan(
		&vehicleNotification.Lat,
		&vehicleNotification.Long,
		&vehicleNotification.Name,
		&vehicleNotification.Phone,
		&vehicleNotification.Type,
		&vehicleNotification.Updated_Description,
	)
	if err != nil {
		log.Error().Err(err).Int("emergency_id", emergencyID).Msg("Failed to get emergency details")
		return
	}

	// Get dispatched vehicles
	rows, err := DB.Query(
		`SELECT vehicle_id FROM dispatched_vehicles WHERE emergency_id = $1`,
		emergencyID,
	)
	if err != nil {
		log.Error().Err(err).Int("emergency_id", emergencyID).Msg("Failed to get dispatched vehicles")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var vehicleID int
		if err := rows.Scan(&vehicleID); err != nil {
			log.Error().Err(err).Msg("Failed to scan vehicle ID")
			continue
		}

		// Update vehicle status to unavailable
		_, err = DB.Exec(
			`UPDATE vehicle_data SET status = false WHERE id = $1`,
			vehicleID,
		)
		if err != nil {
			log.Error().Err(err).Int("vehicle_id", vehicleID).Msg("Failed to update vehicle status")
		}

		// Get vehicle location
		err = DB.QueryRow(
			`SELECT lat, long FROM vehicle_data WHERE id = $1`,
			vehicleID,
		).Scan(&vehicleNotification.Vehicle_Lat, &vehicleNotification.Vehicle_Long)
		if err != nil {
			log.Error().Err(err).Int("vehicle_id", vehicleID).Msg("Failed to get vehicle location")
			continue
		}

		// Get vehicle token
		err = DB.QueryRow(
			`SELECT token FROM vehicle_token_data WHERE vehicle_id = $1`,
			vehicleID,
		).Scan(&vehicleNotification.Token)
		if err != nil {
			log.Error().Err(err).Int("vehicle_id", vehicleID).Msg("Failed to get vehicle token")
			continue
		}

		// Send FCM notification asynchronously
		go VehicleFCM(vehicleNotification)

		log.Info().
			Int("vehicle_id", vehicleID).
			Int("emergency_id", emergencyID).
			Msg("Notification sent to vehicle")
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating vehicle rows")
	}
}
