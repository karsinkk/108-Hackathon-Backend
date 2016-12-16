package usercontroller

import (
	"encoding/json"
	// "fmt"
	"github.com/karsinkk/108/helpers"
	"net/http"
)

func Emergency(res http.ResponseWriter, req *http.Request) {
	var u helpers.EmergencyUserData
	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	// str := fmt.Sprintf("%+v \n", u)
	// fmt.Print(str)

	vehicles_data := make([]helpers.VehicleData, 0)
	vehicles_data = helpers.GetClosest(u.Lat, u.Long, u.Type, u.Number)
	helpers.AddEmergency(vehicles_data, u)
	res.Header().Set("Access-Control-Allow-Origin", "*")

	res.Header().Set("Chiron", "An NP-Incomplete Project")
	res.WriteHeader(200)
	// str := fmt.Sprintf("%+v \n", vehicles_data)
	// fmt.Print(str)
	json.NewEncoder(res).Encode(vehicles_data)
}
