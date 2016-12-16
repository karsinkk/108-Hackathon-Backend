package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/Chiron-Backend/helpers"
	"net/http"
)

func Notification(res http.ResponseWriter, req *http.Request) {

	auth := req.URL.Query().Get("auth")

	Query := fmt.Sprintf("select DISTINCT t.id, t.driver, t.vehicle_no, t.name, t.phone from(select ambulance.id,ambulance.driver, ambulance.vehicle_no, emergency.name, emergency.phone from ambulance join emergency on ambulance.id=emergency.ambulance_id join hospital on ambulance.hosp_id=hospital.id where ambulance.status=false and hospital.auth='%s') as t ", auth)

	rows, _ := helpers.DB.Query(Query)
	defer rows.Close()
	notification := helpers.Notification{}
	notifications := make([]helpers.Notification, 0)
	for rows.Next() {

		if err := rows.Scan(&notification.Amb_id, &notification.Driver, &notification.Vehicle_no, &notification.Name, &notification.Phone); err != nil {
			fmt.Println(err)
		}
		notifications = append(notifications, notification)
	}

	res.Header().Set("Chiron", "An NP-Incomplete Project")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(200)
	str := fmt.Sprintf("%+v \n", notifications)
	fmt.Print(str)
	json.NewEncoder(res).Encode(notifications)
}
