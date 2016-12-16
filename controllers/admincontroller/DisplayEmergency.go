package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/Chiron-Backend/helpers"
	"net/http"
)

func DisplayEmergency(res http.ResponseWriter, req *http.Request) {

	auth := req.URL.Query().Get("auth")

	Query := fmt.Sprintf("select emergency.id,emergency.hosp_id,emergency.ambulance_id,emergency.lat,emergency.long,emergency.name,emergency.phone,emergency.status,emergency.time from emergency,hospital where hospital.id=emergency.hosp_id and hospital.auth='%s'", auth)
	fmt.Println(Query)
	rows, _ := helpers.DB.Query(Query)
	defer rows.Close()
	emergency := helpers.Emergencies{}
	emergencies := make([]helpers.Emergencies, 0)
	for rows.Next() {
		// &emergency.Id, &emergency.Hosp_id, &emergency.Ambulance_id, &emergency.Lat, &emergency.Long, &emergency.Name, &emergency.Phone, &emergency.Status, &emergency.Time)
		if err := rows.Scan(&emergency); err != nil {
			fmt.Println(err)
		}
		emergencies = append(emergencies, emergency)
	}

	res.Header().Set("Chiron", "An NP-Incomplete Project")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(200)
	// str := fmt.Sprintf("%+v \n", emergencies)
	// fmt.Print(str)
	json.NewEncoder(res).Encode(emergencies)
}
