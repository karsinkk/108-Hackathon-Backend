package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108/dif"
	"net/http"
	"strconv"
)

var BaseUrl string = "https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins="

func GetClosest(LatString string, LongString string, Type int, Number int) []VehicleData {
	Conf := dif.ReadConf()
	DB := dif.ConnectDB()
	defer DB.Close()
	list := make(map[int]int, 0)
	var count int = 0
	var Min MIN
	Min.Value = 1<<31 - 1

	Lat, _ := strconv.ParseFloat(LatString, 64)
	Long, _ := strconv.ParseFloat(LongString, 64)

	Query := fmt.Sprintf("select base_id,id, lat,long,phone,driver,vehicle_no,SQRT(POW(69.1 * (cast(lat as float) - %f),2)+POW(69.1*(%f - cast(long as float))*COS(cast(lat as float)/57.3),2)) as distance from vehicles where status=true and type=%d order by distance limit %d", Lat, Long, Type, Number+3)
	rows, _ := DB.Query(Query)
	defer rows.Close()
	vehicle := VehicleData{}
	vehicles := make([]VehicleData, 0)
	for rows.Next() {

		if err := rows.Scan(&vehicle.Base_Id, &vehicle.Id, &vehicle.Lat, &vehicle.Long, &vehicle.Phone, &vehicle.Driver, &vehicle.Vehicle_no, &vehicle.Distance); err != nil {
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
	// fmt.Println(vehicledata)
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
	// fmt.Println(list)
	new_list := SortMapByValue(list)
	// fmt.Println("SOrted")
	// fmt.Println(new_list)

	vehicles_return := make([]VehicleData, 0)
	for _, v := range new_list {
		vehicles[v.Key].Time = v.Value
		vehicles_return = append(vehicles_return, vehicles[v.Key])
	}
	// str := fmt.Sprintf("%+v", vehicles_return)
	// fmt.Println(str)
	return vehicles_return[0:Number]
}
