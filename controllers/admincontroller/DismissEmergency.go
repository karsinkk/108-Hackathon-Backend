package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108/dif"
	"github.com/karsinkk/108/helpers"
	"net/http"
)

func DismissEmergency(res http.ResponseWriter, req *http.Request) {
	DB := dif.ConnectDB()
	var r helpers.DismissEmergencyData
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	Query := fmt.Sprintf("update emergency set status=false,dismissed=true,updated_description='%s' where id='%d'", r.Emergency_Id, r.Dismissed_Reason)
	_, erro := DB.Query(Query)
	if err != nil {
		fmt.Println(erro)
	}
	Query = fmt.Sprintf("delete from dispatched_vehicles where emergency_id='%d'", r.Emergency_Id)
	_, erro = DB.Query(Query)
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
