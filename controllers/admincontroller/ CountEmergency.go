package admincontroller

import (
	"context"
	"net/http"
	"time"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// Count represents the emergency count response
type Count struct {
	Id int `json:"id"`
}

// CountEmergency handles WebSocket connections for emergency count updates
func CountEmergency(res http.ResponseWriter, req *http.Request) {
	upgrader := helpers.GetUpgrader()
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to upgrade WebSocket connection")
		return
	}
	defer conn.Close()

	DB := dif.GetDB()

	// Send initial count
	var countData Count
	err = DB.QueryRow(`SELECT COUNT(id) FROM emergency WHERE seen = false`).Scan(&countData.Id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to query emergency count")
		return
	}

	if err := conn.WriteJSON(countData); err != nil {
		log.Error().Err(err).Msg("Failed to write initial count to WebSocket")
		return
	}

	// Create context for cleanup
	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	// Periodic updates
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Debug().Msg("Emergency count WebSocket connection closed")
			return
		case <-ticker.C:
			countData = Count{}
			err := DB.QueryRow(`SELECT COUNT(id) FROM emergency WHERE seen = false`).Scan(&countData.Id)
			if err != nil {
				log.Error().Err(err).Msg("Failed to query emergency count")
				continue
			}

			if err := conn.WriteJSON(countData); err != nil {
				log.Error().Err(err).Msg("Failed to write count to WebSocket")
				return
			}
		}
	}
}
