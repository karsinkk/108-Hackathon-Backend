package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/rs/zerolog/log"
)

// FCMVehicleData represents the FCM payload for vehicle notifications
type FCMVehicleData struct {
	Lat                string `json:"Lat"`
	Long               string `json:"Long"`
	VehicleLat         string `json:"Vehicle_Lat"`
	VehicleLong        string `json:"Vehicle_Long"`
	Name               string `json:"Name"`
	Phone              string `json:"Phone"`
	UpdatedDescription string `json:"Updated_Description"`
}

// FCMVehicleMessage represents the FCM message structure
type FCMVehicleMessage struct {
	Data FCMVehicleData `json:"data"`
	To   string         `json:"to"`
}

// VehicleFCM sends a Firebase Cloud Messaging notification to a vehicle
func VehicleFCM(vehicleNotification VehicleNotificationData) {
	config := dif.GetConfig()
	fcmKey := "key=" + config.VehicleFCMKey

	// Build FCM payload using proper JSON marshaling
	message := FCMVehicleMessage{
		Data: FCMVehicleData{
			Lat:                vehicleNotification.Lat,
			Long:               vehicleNotification.Long,
			VehicleLat:         vehicleNotification.Vehicle_Lat,
			VehicleLong:        vehicleNotification.Vehicle_Long,
			Name:               vehicleNotification.Name,
			Phone:              vehicleNotification.Phone,
			UpdatedDescription: vehicleNotification.Updated_Description,
		},
		To: vehicleNotification.Token,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal FCM message")
		return
	}

	url := "https://fcm.googleapis.com/fcm/send"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create FCM request")
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fcmKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send FCM notification")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Warn().Int("status", resp.StatusCode).Msg("FCM notification returned non-200 status")
	} else {
		log.Debug().Msg("Vehicle FCM notification sent successfully")
	}
}
