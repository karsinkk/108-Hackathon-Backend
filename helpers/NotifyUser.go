package helpers

import (
	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/rs/zerolog/log"
)

// NotifyUser sends notification to the user about their emergency response
func NotifyUser(emergencyID int) {
	DB := dif.GetDB()

	// Get the closest vehicle details
	var vehicleID int
	userNotification := UserNotificationData{}

	err := DB.QueryRow(
		`SELECT vehicle_id, time_taken
		 FROM dispatched_vehicles
		 WHERE emergency_id = $1
		 ORDER BY time_taken
		 LIMIT 1`,
		emergencyID,
	).Scan(&vehicleID, &userNotification.Time)

	if err != nil {
		log.Error().Err(err).Int("emergency_id", emergencyID).Msg("Failed to get dispatched vehicle")
		return
	}

	// Convert seconds to minutes
	userNotification.Time = userNotification.Time / 60

	// Get vehicle details
	err = DB.QueryRow(
		`SELECT lat, long, vehicle_no, phone, driver
		 FROM vehicle_data
		 WHERE id = $1`,
		vehicleID,
	).Scan(
		&userNotification.Lat,
		&userNotification.Long,
		&userNotification.Vehicle_No,
		&userNotification.Phone,
		&userNotification.Name,
	)

	if err != nil {
		log.Error().Err(err).Int("vehicle_id", vehicleID).Msg("Failed to get vehicle details")
		return
	}

	// Get user's FCM token
	err = DB.QueryRow(
		`SELECT token FROM emergency_token_data WHERE emergency_id = $1`,
		emergencyID,
	).Scan(&userNotification.Token)

	if err != nil {
		log.Error().Err(err).Int("emergency_id", emergencyID).Msg("Failed to get user token")
		return
	}

	// Send FCM notification
	UserFCM(userNotification)

	log.Info().
		Int("emergency_id", emergencyID).
		Int("vehicle_id", vehicleID).
		Msg("User notification sent")
}
