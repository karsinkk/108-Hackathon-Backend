package admincontroller

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

const notificationQuery = `
	SELECT
		emergency.id, emergency.lat, emergency.long, emergency.phone, emergency.name,
		emergency.time, emergency.status, emergency.type, emergency.description,
		emergency.seen, emergency.updated_time, d.vehicle_id, d.time_taken,
		v.district, v.name, v.phone, v.lat, v.long, v.driver, v.vehicle_no
	FROM emergency
	JOIN dispatched_vehicles d ON emergency.id = d.emergency_id
	JOIN vehicle_data v ON v.id = d.vehicle_id
	WHERE emergency.status = TRUE AND v.status = TRUE
	ORDER BY emergency.id`

// sendNotifications sends periodic notification updates via WebSocket
func sendNotifications(ctx context.Context, conn *websocket.Conn) {
	DB := dif.GetDB()

	// Send initial data
	notifications := queryNotifications(DB)
	if err := conn.WriteJSON(notifications); err != nil {
		log.Error().Err(err).Msg("Failed to write notifications to WebSocket")
		return
	}

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Debug().Msg("Notification WebSocket connection closed")
			return
		case <-ticker.C:
			notifications := queryNotifications(DB)
			if err := conn.WriteJSON(notifications); err != nil {
				log.Error().Err(err).Msg("Failed to write notifications to WebSocket")
				return
			}
		}
	}
}

// queryNotifications fetches notification data from the database
func queryNotifications(DB interface {
	Query(query string, args ...interface{}) (interface {
		Close() error
		Next() bool
		Scan(dest ...interface{}) error
		Err() error
	}, error)
}) []helpers.Notification {
	rows, err := dif.GetDB().Query(notificationQuery)
	if err != nil {
		log.Error().Err(err).Msg("Failed to query notifications")
		return []helpers.Notification{}
	}
	defer rows.Close()

	notifications := make([]helpers.Notification, 0)
	for rows.Next() {
		var n helpers.Notification
		if err := rows.Scan(
			&n.Eid, &n.ELat, &n.ELong, &n.Phone_1, &n.Name_1,
			&n.Time, &n.Status, &n.Type, &n.Description,
			&n.Seen, &n.Updated_time, &n.Vehicle_id, &n.Time_taken,
			&n.District, &n.Name_2, &n.Phone_2, &n.VLat, &n.VLong,
			&n.Driver, &n.Vehicle_no,
		); err != nil {
			log.Error().Err(err).Msg("Failed to scan notification row")
			continue
		}
		notifications = append(notifications, n)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating notification rows")
	}

	return notifications
}

// Notification handles WebSocket connections for notification updates
func Notification(res http.ResponseWriter, req *http.Request) {
	upgrader := helpers.GetUpgrader()
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to upgrade WebSocket connection")
		return
	}

	ctx, cancel := context.WithCancel(req.Context())

	go func() {
		defer cancel()
		defer conn.Close()

		var baseID helpers.AdminNotificationPostData
		if err := conn.ReadJSON(&baseID); err != nil {
			log.Error().Err(err).Msg("Failed to read base ID from WebSocket")
			return
		}

		log.Debug().Int("base_id", baseID.Id).Msg("Starting notification stream")
		sendNotifications(ctx, conn)
	}()
}
