package vehiclecontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"net/http"
)

func LoginVehicle(res http.ResponseWriter, req *http.Request) {
	DB := dif.ConnectDB()
	defer DB.Close()

	var a helpers.VehicleLoginData
	_ = json.NewDecoder(req.Body).Decode(&a)
	str := fmt.Sprintf("%+v \n", a)
	fmt.Print(str)
	logindata := helpers.LoginData{Username: a.Username, Password: a.Password}
	auth := helpers.LoginUser(logindata)
	fmt.Println(auth)
	var v helpers.Vehicle_Id
	if auth != "" {
		Query := fmt.Sprintf("select id from vehicle_data where username='%s'", a.Username)
		fmt.Println(Query)
		row := DB.QueryRow(Query)
		_ = row.Scan(&v.Id)
		Query = fmt.Sprintf("update vehicle_data set lat='%s',long='%s',time=now()::timestamp where id=%d ", a.Lat, a.Long, v.Id)
		fmt.Println(Query)
		_ = DB.QueryRow(Query)
		Query = fmt.Sprintf("update vehicle_token_data set token='%s' where id='%d'", a.Token, v.Id)
		fmt.Println(Query)
		_ = DB.QueryRow(Query)
	}
	res.Header().Set("Access-Control-Allow-Origin", "*")

	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(200)
	str = fmt.Sprintf("%+v \n", v)
	fmt.Print(str)
	json.NewEncoder(res).Encode(v)
}
