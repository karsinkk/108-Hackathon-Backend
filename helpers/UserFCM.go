package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/rs/zerolog/log"
)

// FCMUserData represents the FCM payload for user notifications
type FCMUserData struct {
	Lat       string `json:"Lat"`
	Long      string `json:"Long"`
	Time      int    `json:"Time"`
	VehicleNo string `json:"Vehicle_no"`
	Phone     string `json:"Phone"`
	Name      string `json:"Name"`
}

// FCMUserMessage represents the FCM message structure
type FCMUserMessage struct {
	Data FCMUserData `json:"data"`
	To   string      `json:"to"`
}

// UserFCM sends a Firebase Cloud Messaging notification to a user
func UserFCM(userNotification UserNotificationData) {
	config := dif.GetConfig()
	fcmKey := "key=" + config.UserFCMKey

	// Build FCM payload using proper JSON marshaling
	message := FCMUserMessage{
		Data: FCMUserData{
			Lat:       userNotification.Lat,
			Long:      userNotification.Long,
			Time:      userNotification.Time,
			VehicleNo: userNotification.Vehicle_No,
			Phone:     userNotification.Phone,
			Name:      userNotification.Name,
		},
		To: userNotification.Token,
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
		log.Debug().Msg("User FCM notification sent successfully")
	}
}
