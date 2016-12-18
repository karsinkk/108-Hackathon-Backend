package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108/dif"
	"github.com/karsinkk/108/helpers"
	"net/http"
)

func AddVehicle(res http.ResponseWriter, req *http.Request) {
	DB := dif.ConnectDB()
	var r helpers.VehicleAddData
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	str := fmt.Sprintf("%+v \n", r)
	fmt.Print(str)

	Query := fmt.Sprintf("insert into vehicle_data(vehicle_no,driver,phone,type) values('%s','%s','%s','1')", r.Vehicle_no, r.Driver, r.Phone)
	_ = DB.QueryRow(Query)

	var Id int
	Query = fmt.Sprintf("select id from vehicle_data where vehicle_no='%s'", r.Vehicle_no)
	row := DB.QueryRow(Query)
	_ = row.Scan(&Id)

	username := fmt.Sprintf("vehicle%d", Id)

	Query = fmt.Sprintf("insert into vehicle_data(username) values('%s') where id='%d'", username, Id)
	_ = DB.QueryRow(Query)
	data := helpers.AdminRegisterData{Username: username}
	_ = helpers.RegisterUser(data)
	res.Header().Set("108", "An NP-Incomplete Project")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(200)
	fmt.Fprint(res, username)

}
