package vehiclecontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/Chiron-Backend/helpers"
	"net/http"
)

type Loc struct {
	Lat  string
	Long string
}

func UpdateVehicle(res http.ResponseWriter, req *http.Request) {
	var a helpers.AmbUpdateData
	err := json.NewDecoder(req.Body).Decode(&a)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	st := fmt.Sprintf("%+v \n", a)
	fmt.Print(st)
	add_query := ""
	if a.Phone != "" {
		add_query += ",phone='" + a.Phone + "'"
	}
	if a.Driver != "" {
		add_query += ",driver='" + a.Driver + "'"
	}

	Query := fmt.Sprintf("update ambulance set lat='%s', long='%s', time=now()::timestamp,status=%t %s where id=%d", a.Lat, a.Long, a.Status, add_query, a.Id)
	// fmt.Println(Query)

	_, _ = helpers.DB.Query(Query)
	var flag bool
	loc := Loc{}
	Query = fmt.Sprintf("select status from ambulance")
	_ = helpers.DB.QueryRow(Query).Scan(&flag)
	if flag == false {
		Query = fmt.Sprintf("select emergency.lat, emergency.long from emergency,ambulance where emergency.ambulance_id=%d", a.Id)
		_ = helpers.DB.QueryRow(Query).Scan(&loc.Lat, &loc.Long)
		res.Header().Set("Chiron", "An NP-Incomplete Project")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.WriteHeader(200)
		str := fmt.Sprintf("%+v \n", a.Id)
		fmt.Print(str)
		json.NewEncoder(res).Encode(loc)
	}

}
