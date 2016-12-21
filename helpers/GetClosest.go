package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"net/http"
	"strconv"
)

var BaseUrl string = "https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins="

func GetClosest(LatString string, LongString string, Type int, N int) []VehicleData {
	Conf := dif.ReadConf()
	DB := dif.ConnectDB()
	defer DB.Close()
	list := make(map[int]int, 0)
	var count int = 0
	var Min MIN
	Min.Value = 1<<31 - 1

	var Number int
	Number = N/3 + 1
	if Number > 7 {
		Number = 7
	}

	Lat, _ := strconv.ParseFloat(LatString, 64)
	Long, _ := strconv.ParseFloat(LongString, 64)
	fmt.Println(Lat, Long)
	Query := fmt.Sprintf("select id, lat,long,SQRT(POW(69.1 * (cast(lat as float) - %f),2)+POW(69.1*(%f - cast(long as float))*COS(cast(lat as float)/57.3),2)) as distance from vehicle_data where status=true order by distance limit %d", Lat, Long, Number+3)
	fmt.Println(Query)
	rows, _ := DB.Query(Query)
	defer rows.Close()
	vehicle := VehicleData{}
	vehicles := make([]VehicleData, 0)
	for rows.Next() {

		if err := rows.Scan(&vehicle.Id, &vehicle.Lat, &vehicle.Long, &vehicle.Distance); err != nil {
			fmt.Println(err)
		}
		vehicles = append(vehicles, vehicle)
	}
	vehicledata := ""
	for k, v := range vehicles {
		if k > 0 {
			vehicledata += "|"
		}
		vehicledata += fmt.Sprintf("%s,%s", v.Lat, v.Long)
	}
	fmt.Println(vehicledata)
	url := fmt.Sprintf("%s%s,%s&destinations=%s&key=%s", BaseUrl, LatString, LongString, vehicledata, Conf.API_KEY)

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
			list[count] = j.Duration.Value
			count++
		}
	}
	fmt.Println(list)
	fmt.Println("SOrted")

	new_list := SortMapByValue(list)
	fmt.Println(new_list)

	vehicles_return := make([]VehicleData, 0)
	for _, v := range new_list {
		vehicles[v.Key].Time = v.Value
		vehicles_return = append(vehicles_return, vehicles[v.Key])
	}
	str := fmt.Sprintf("%+v", vehicles_return)
	fmt.Println(str)
	return vehicles_return[0:Number]
}
