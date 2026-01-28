package admincontroller

import (
	"context"
	"net/http"
	"time"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// DisplayEmergency handles WebSocket connections for emergency data streaming
func DisplayEmergency(res http.ResponseWriter, req *http.Request) {
	upgrader := helpers.GetUpgrader()
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to upgrade WebSocket connection")
		return
	}
	defer conn.Close()

	DB := dif.GetDB()

	// Send initial data
	emergencies := queryEmergencies(DB)
	if err := conn.WriteJSON(emergencies); err != nil {
		log.Error().Err(err).Msg("Failed to write emergencies to WebSocket")
		return
	}

	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Debug().Msg("Emergency display WebSocket connection closed")
			return
		case <-ticker.C:
			emergencies := queryEmergencies(DB)
			if err := conn.WriteJSON(emergencies); err != nil {
				log.Error().Err(err).Msg("Failed to write emergencies to WebSocket")
				return
			}
		}
	}
}

// queryEmergencies fetches all emergency data from the database
func queryEmergencies(DB interface {
	Query(query string, args ...interface{}) (interface {
		Close() error
		Next() bool
		Scan(dest ...interface{}) error
		Err() error
	}, error)
}) []helpers.EmergencyData {
	rows, err := dif.GetDB().Query(`SELECT * FROM emergency`)
	if err != nil {
		log.Error().Err(err).Msg("Failed to query emergencies")
		return []helpers.EmergencyData{}
	}
	defer rows.Close()

	emergencies := make([]helpers.EmergencyData, 0)
	for rows.Next() {
		var e helpers.EmergencyData
		if err := rows.Scan(
			&e.Id, &e.Lat, &e.Long, &e.Phone, &e.Name,
			&e.Status, &e.Time, &e.Type, &e.Description,
			&e.Seen, &e.Updated_time, &e.Updated_description, &e.Dismissed,
		); err != nil {
			log.Error().Err(err).Msg("Failed to scan emergency row")
			continue
		}
		emergencies = append(emergencies, e)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating emergency rows")
	}

	return emergencies
}
