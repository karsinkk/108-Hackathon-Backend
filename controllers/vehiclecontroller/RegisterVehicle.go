package vehiclecontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/Chiron-Backend/helpers"
	"net/http"
)

func RegisterVehicle(res http.ResponseWriter, req *http.Request) {
	var a helpers.AmbRegisterToken
	_ = json.NewDecoder(req.Body).Decode(&a)

	Query := fmt.Sprintf("select ambulance.id, vehicle_no from ambulance,hospital where ambulance.hosp_id = hospital.id and hospital.auth_token = '%s'", a.Auth_Token)
	// fmt.Println(Query)

	rows, _ := helpers.DB.Query(Query)
	defer rows.Close()
	amb_register_data := helpers.AmbRegisterData{}
	amb_register_data_list := make([]helpers.AmbRegisterData, 0)
	for rows.Next() {

		if err := rows.Scan(&amb_register_data.Id, &amb_register_data.Vehicle_no); err != nil {
			fmt.Println(err)
		}
		amb_register_data_list = append(amb_register_data_list, amb_register_data)
	}
	res.Header().Set("Access-Control-Allow-Origin", "*")

	res.Header().Set("Chiron", "An NP-Incomplete Project")
	res.WriteHeader(200)
	str := fmt.Sprintf("%+v \n", amb_register_data_list)
	fmt.Print(str)
	json.NewEncoder(res).Encode(amb_register_data_list)
}
