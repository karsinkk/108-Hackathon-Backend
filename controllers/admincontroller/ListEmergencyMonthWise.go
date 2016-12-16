package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/Chiron-Backend/helpers"
	"net/http"
)

func ListEmergencyMonthWise(res http.ResponseWriter, req *http.Request) {

	auth := req.URL.Query().Get("auth")
	add_query := ""
	if auth != "" {
		add_query = "where hospital.auth='" + auth + "'"
	}
	Query := fmt.Sprintf("select to_char(time,'Mon') as mon, extract(year from time) as yyyy, count(emergency.id) as ID from emergency join hospital on hospital.id=emergency.hosp_id %s group by 1,2", add_query)
	// fmt.Println(Query)

	rows, _ := helpers.DB.Query(Query)
	defer rows.Close()
	emergency_monthwise := helpers.EmergencyMonthWise{}
	emergencies_monthwise := make([]helpers.EmergencyMonthWise, 0)
	for rows.Next() {

		if err := rows.Scan(&emergency_monthwise.Month, &emergency_monthwise.Year, &emergency_monthwise.Id); err != nil {
			fmt.Println(err)
		}
		emergencies_monthwise = append(emergencies_monthwise, emergency_monthwise)
	}
	res.Header().Set("Access-Control-Allow-Origin", "*")

	res.Header().Set("Chiron", "An NP-Incomplete Project")
	res.WriteHeader(200)
	// str := fmt.Sprintf("%+v \n", emergencies_monthwise)
	// 	fmt.Print(str)
	json.NewEncoder(res).Encode(emergencies_monthwise)
}
