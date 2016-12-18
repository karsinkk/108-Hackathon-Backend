package vehiclecontroller

import (
	// "encoding/json"
	"fmt"
	"github.com/karsinkk/108/dif"
	"github.com/karsinkk/108/helpers"
	"net/http"
)

type Loc struct {
	Lat  string
	Long string
}

func UpdateVehicle(res http.ResponseWriter, req *http.Request) {

	DB := dif.ConnectDB()
	conn, err := helpers.Upgrader.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for {
			var update_data helpers.VehicleUpdateData
			err := conn.ReadJSON(&update_data)

			if err != nil {
				fmt.Println("IN error")
				fmt.Println(err)
				conn.Close()
				return
			}

			str := fmt.Sprintf("%+v", update_data)
			fmt.Println("In sprintf", str)

			Query := fmt.Sprintf("update vehicle_data set lat='%s',long='%s'where id='%d'", update_data.Lat, update_data.Long, update_data.Id)
			fmt.Println(Query)
			_ = DB.QueryRow(Query)
			Query = fmt.Sprintf("update vehicle_token_data set token='%s' where id='%d'", update_data.Token, update_data.Id)
			fmt.Println(Query)
			_ = DB.QueryRow(Query)
		}
	}()

}
