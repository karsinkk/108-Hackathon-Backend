package vehiclecontroller

import (
	"context"
	"net/http"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"github.com/rs/zerolog/log"
)

// UpdateVehicle handles WebSocket connections for real-time vehicle location updates
func UpdateVehicle(res http.ResponseWriter, req *http.Request) {
	upgrader := helpers.GetUpgrader()
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to upgrade WebSocket connection for vehicle update")
		return
	}

	DB := dif.GetDB()

	ctx, cancel := context.WithCancel(req.Context())

	go func() {
		defer cancel()
		defer conn.Close()

		for {
			select {
			case <-ctx.Done():
				log.Debug().Msg("Vehicle update WebSocket connection closed")
				return
			default:
				var updateData helpers.VehicleUpdateData
				if err := conn.ReadJSON(&updateData); err != nil {
					log.Debug().Err(err).Msg("WebSocket read error, closing connection")
					return
				}

				log.Debug().
					Int("vehicle_id", updateData.Id).
					Str("lat", updateData.Lat).
					Str("long", updateData.Long).
					Msg("Updating vehicle location")

				// Update vehicle location using parameterized query
				_, err := DB.Exec(
					`UPDATE vehicle_data SET lat = $1, long = $2 WHERE id = $3`,
					updateData.Lat, updateData.Long, updateData.Id,
				)
				if err != nil {
					log.Error().Err(err).Int("vehicle_id", updateData.Id).Msg("Failed to update vehicle location")
				}

				// Update vehicle token using parameterized query
				_, err = DB.Exec(
					`UPDATE vehicle_token_data SET token = $1 WHERE vehicle_id = $2`,
					updateData.Token, updateData.Id,
				)
				if err != nil {
					log.Error().Err(err).Int("vehicle_id", updateData.Id).Msg("Failed to update vehicle token")
				}
			}
		}
	}()
}
