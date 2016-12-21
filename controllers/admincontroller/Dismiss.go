package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"net/http"
)

func Dismiss(res http.ResponseWriter, req *http.Request) {
	DB := dif.ConnectDB()
	var r helpers.DismissData
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	Query := fmt.Sprintf("delete from dispatched_vehicles where vehicle_id=%d and emergency_id=%d", r.Vehicle_Id, r.Emergency_Id)
	_, erro := DB.Query(Query)
	if err != nil {
		fmt.Println(erro)
	}
	// fmt.Println(Query)
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(200)
	fmt.Fprintf(res, "1")
	// str := fmt.Sprintf("%+v \n", r)
	// fmt.Print(str)
}
