package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/Chiron-Backend/helpers"
	"net/http"
)

func CountAmbulances(res http.ResponseWriter, req *http.Request) {

	auth := req.URL.Query().Get("auth")

	Query := fmt.Sprintf("select count(hospital.id) from ambulance JOIN hospital on hospital.id=ambulance.hosp_id where hospital.auth='%s' and ambulance.status=true UNION select count(hospital.id) count2 from ambulance JOIN hospital on hospital.id=ambulance.hosp_id where hospital.auth='%s' and ambulance.status=false", auth, auth)
	// fmt.Println(Query)

	rows, _ := helpers.DB.Query(Query)
	defer rows.Close()
	amb_count := helpers.Ambulance_Count{}

	rows.Next()
	rows.Scan(&amb_count.OffDuty)
	rows.Next()
	rows.Scan(&amb_count.OnDuty)
	res.Header().Set("Access-Control-Allow-Origin", "*")

	res.Header().Set("Chiron", "An NP-Incomplete Project")
	res.WriteHeader(200)
	// str := fmt.Sprintf("%+v \n", amb_count)
	// fmt.Print(str)
	json.NewEncoder(res).Encode(amb_count)
}
