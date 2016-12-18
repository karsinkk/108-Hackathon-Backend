package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108/dif"
	"github.com/karsinkk/108/helpers"
	"net/http"
)

func Status(res http.ResponseWriter, req *http.Request) {
	DB := dif.ConnectDB()
	var s helpers.StatusData
	err := json.NewDecoder(req.Body).Decode(&s)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	str := fmt.Sprintf("%+v \n", s)
	fmt.Print(str)
	Query := fmt.Sprintf("update emergency set dismissed='%v',updated_description='%s' where id='%d'", s.Dispatched, s.Updated_Description, s.Emergency_Id)
	fmt.Println(Query)
	_ = DB.QueryRow(Query)
	if s.Dispatched == true {
		helpers.NotifyVehicle(s.Emergency_Id, s.User_Id)
		helpers.NotifyUser(s.Emergency_Id)
	}
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(200)
}
