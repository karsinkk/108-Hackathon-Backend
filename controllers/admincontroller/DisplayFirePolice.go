package admincontroller

import (
	"context"
	"net/http"
	"time"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// DisplayFirePolice handles WebSocket connections for fire/police emergency updates
func DisplayFirePolice(res http.ResponseWriter, req *http.Request) {
	upgrader := helpers.GetUpgrader()
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to upgrade WebSocket connection")
		return
	}
	defer conn.Close()

	DB := dif.GetDB()

	// Query for non-ambulance emergencies (type != 1)
	query := `SELECT * FROM emergency WHERE type != 1 AND status = true`

	// Send initial data
	emergencies := queryFirePoliceEmergencies(DB, query)
	if err := conn.WriteJSON(emergencies); err != nil {
		log.Error().Err(err).Msg("Failed to write fire/police emergencies to WebSocket")
		return
	}

	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Debug().Msg("Fire/Police WebSocket connection closed")
			return
		case <-ticker.C:
			emergencies := queryFirePoliceEmergencies(DB, query)
			if err := conn.WriteJSON(emergencies); err != nil {
				log.Error().Err(err).Msg("Failed to write fire/police emergencies to WebSocket")
				return
			}
		}
	}
}

// queryFirePoliceEmergencies fetches fire/police emergency data
func queryFirePoliceEmergencies(DB interface {
	Query(query string, args ...interface{}) (interface {
		Close() error
		Next() bool
		Scan(dest ...interface{}) error
		Err() error
	}, error)
}, query string) []helpers.EmergencyData {
	rows, err := dif.GetDB().Query(query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to query fire/police emergencies")
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
