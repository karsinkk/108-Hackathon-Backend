package helpers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/rs/zerolog/log"
)

var baseURL = "https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins="

// GetClosest finds the closest available vehicles to a given location
func GetClosest(latString string, longString string, vehicleType int, n int) []VehicleData {
	config := dif.GetConfig()
	DB := dif.GetDB()

	list := make(map[int]int, 0)
	count := 0

	number := n/3 + 1
	if number > 7 {
		number = 7
	}

	lat, err := strconv.ParseFloat(latString, 64)
	if err != nil {
		log.Error().Err(err).Str("lat", latString).Msg("Failed to parse latitude")
		return nil
	}

	long, err := strconv.ParseFloat(longString, 64)
	if err != nil {
		log.Error().Err(err).Str("long", longString).Msg("Failed to parse longitude")
		return nil
	}

	log.Debug().Float64("lat", lat).Float64("long", long).Msg("Finding closest vehicles")

	// Use parameterized query for distance calculation
	rows, err := DB.Query(
		`SELECT id, lat, long,
		 SQRT(POW(69.1 * (CAST(lat AS FLOAT) - $1), 2) +
		      POW(69.1 * ($2 - CAST(long AS FLOAT)) * COS(CAST(lat AS FLOAT) / 57.3), 2)) AS distance
		 FROM vehicle_data
		 WHERE status = true
		 ORDER BY distance
		 LIMIT $3`,
		lat, long, number+3,
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to query vehicles")
		return nil
	}
	defer rows.Close()

	vehicles := make([]VehicleData, 0)
	for rows.Next() {
		var vehicle VehicleData
		if err := rows.Scan(&vehicle.Id, &vehicle.Lat, &vehicle.Long, &vehicle.Distance); err != nil {
			log.Error().Err(err).Msg("Failed to scan vehicle row")
			continue
		}
		vehicles = append(vehicles, vehicle)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating vehicle rows")
		return nil
	}

	if len(vehicles) == 0 {
		log.Warn().Msg("No available vehicles found")
		return vehicles
	}

	// Build destination coordinates for Google Maps API
	vehicleData := ""
	for k, v := range vehicles {
		if k > 0 {
			vehicleData += "|"
		}
		vehicleData += v.Lat + "," + v.Long
	}

	// Call Google Distance Matrix API
	url := baseURL + latString + "," + longString + "&destinations=" + vehicleData + "&key=" + config.APIKey

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create Google Maps request")
		return vehicles
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to call Google Maps API")
		return vehicles
	}
	defer resp.Body.Close()

	var data DistanceMatrixResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Error().Err(err).Msg("Failed to decode Google Maps response")
		return vehicles
	}

	// Process distance matrix response
	for _, i := range data.Rows {
		for _, j := range i.Elements {
			list[count] = j.Duration.Value
			count++
		}
	}

	// Sort by travel time
	newList := SortMapByValue(list)

	vehiclesReturn := make([]VehicleData, 0)
	for _, v := range newList {
		if v.Key < len(vehicles) {
			vehicles[v.Key].Time = v.Value
			vehiclesReturn = append(vehiclesReturn, vehicles[v.Key])
		}
	}

	log.Info().Int("count", len(vehiclesReturn)).Msg("Found closest vehicles")

	if len(vehiclesReturn) < number {
		return vehiclesReturn
	}
	return vehiclesReturn[0:number]
}
