package helpers

import (
	"fmt"
	"github.com/karsinkk/108/dif"
	"strconv"
)

func AddEmergency(vehicles_data []VehicleData, u EmergencyUserData) {

	DB := dif.ConnectDB()
	defer DB.Close()
	amb_id, _ := strconv.ParseInt(amb.Id, 10, 64)
	Query := fmt.Sprintf("insert into emergency(hosp_id,ambulance_id,lat,long,name,phone) values('%d','%d','%s','%s','%s','%s')", amb.Base_Id, amb_id, userLat, userLong, name, phone)
	_ = DB.QueryRow(Query)

	Query = fmt.Sprintf("update ambulance set status=false where id=%d", amb_id)
	fmt.Println(Query)
	_ = DB.QueryRow(Query)
}
