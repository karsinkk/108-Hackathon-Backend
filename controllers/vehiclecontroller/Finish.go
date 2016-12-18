package vehiclecontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108/dif"
	"github.com/karsinkk/108/helpers"
	"net/http"
)

func Finish(res http.ResponseWriter, req *http.Request) {
	DB := dif.ConnectDB()
	defer DB.Close()

	var a helpers.Vehicle_Id
	_ = json.NewDecoder(req.Body).Decode(&a)
	str := fmt.Sprintf("%+v \n", a)
	fmt.Print(str)
	var E_Id int
	Query := fmt.Sprintf("select emergency_id from dispatched_vehicles where vehicle_id='%d", a.Id)
	row := DB.QueryRow(Query)
	_ = row.Scan(&E_Id)

	Query = fmt.Sprintf("update emergency set status=false where id='%d'", E_Id)
	_ = DB.QueryRow(Query)

	Query = fmt.Sprintf("update vehicle_data set status=true where id='%d'", a.Id)
	_ = DB.QueryRow(Query)

	fmt.Fprint(res, E_Id)
}
