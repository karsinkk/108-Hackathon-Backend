package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/Chiron-Backend/helpers"
	"net/http"
)

func ListAmbulances(res http.ResponseWriter, req *http.Request) {

	auth := req.URL.Query().Get("auth")

	Query := fmt.Sprintf("select ambulance.hosp_id, ambulance.id, ambulance.lat, ambulance.long, ambulance.phone, ambulance.time, ambulance.status, ambulance.driver, ambulance.vehicle_no from ambulance join hospital on hospital.id=ambulance.hosp_id where hospital.auth='%s'", auth)

	rows, _ := helpers.DB.Query(Query)
	defer rows.Close()
	amb := helpers.Ambulance{}
	ambs := make([]helpers.Ambulance, 0)
	for rows.Next() {
		if err := rows.Scan(&amb.Hosp_id, &amb.Id, &amb.Lat, &amb.Long, &amb.Phone, &amb.Time, &amb.Status, &amb.Driver, &amb.Vehicle_no); err != nil {
			fmt.Println(err)
		}
		ambs = append(ambs, amb)
	}

	res.Header().Set("Access-Control-Allow-Origin", "*")

	res.Header().Set("Chiron", "An NP-Incomplete Project")
	res.WriteHeader(200)
	// str := fmt.Sprintf("%+v \n", amb_count)
	// fmt.Print(str)
	json.NewEncoder(res).Encode(ambs)
}
