package helpers

import (
	"fmt"
	"github.com/karsinkk/108-Hackathon-Backend/dif"
)

func AddEmergency(vehicles_data []VehicleData, u EmergencyUserData) {

	DB := dif.ConnectDB()
	defer DB.Close()

	Query := fmt.Sprintf("insert into emergency(lat,long,name,phone,type,description) values('%s','%s','%s','%s','%d','%d')", u.Lat, u.Long, u.Name, u.Phone, u.Type, u.Description)
	fmt.Println(Query)
	_ = DB.QueryRow(Query)
	Query = fmt.Sprintf("select id from emergency order by time  desc limit 1")
	fmt.Println(Query)
	var Id int
	row := DB.QueryRow(Query)
	_ = row.Scan(&Id)
	if u.Type == 1 {
		for _, v := range vehicles_data {
			Query = fmt.Sprintf("insert into dispatched_vehicles(emergency_id,vehicle_id,time_taken,distance) values('%d','%d','%d','%f')", Id, v.Id, v.Time, v.Distance)
			fmt.Println(Query)
			_ = DB.QueryRow(Query)
		}
	}
	Query = fmt.Sprintf("insert into emergency_token_data(emergency_id,token) values('%d','%s')", Id, u.Token)
	fmt.Println(Query)
	_ = DB.QueryRow(Query)

}
