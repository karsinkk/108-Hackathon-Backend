package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Duration struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type Distance struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type DistanceMatrixElement struct {
	Duration Duration `json:"duration"`
	Distance Distance `json:"distance"`
	Status   string   `json:"status"`
}

type DistanceMatrixElementsRow struct {
	Elements []DistanceMatrixElement `json:"elements"`
}

type DistanceMatrixResponse struct {
	OriginAddresses      []string                    `json:"origin_addresses"`
	DestinationAddresses []string                    `json:"destination_addresses"`
	Rows                 []DistanceMatrixElementsRow `json:"rows"`
}

var BaseUrl string = "https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins="

var DurationData = make(map[int]int)

type MIN struct {
	Key   int
	Value int
}

func NearestBase(Lat string, Long string) (ID int, Time int) {
	var count int = 0
	var Min MIN
	Min.Value = 1<<31 - 1
	for _, Base := range Conf.BaseLocationDataEncoded {
		url := fmt.Sprintf("%s%s,%s&destinations=enc:%s:&key=%s", BaseUrl, Lat, Long, Base, Conf.API_KEY)
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var data DistanceMatrixResponse
		json.NewDecoder(resp.Body).Decode(&data)
		for _, i := range data.Rows {
			for _, j := range i.Elements {
				DValue := j.Duration.Value
				if DValue < Min.Value {
					Min.Key = count
					Min.Value = DValue
				}
				count++
			}
		}
	}
	return Min.Key, Min.Value
}
