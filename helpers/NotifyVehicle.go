package helpers

import (
	"fmt"
	"github.com/karsinkk/108-Hackathon-Backend/dif"
)

func NotifyVehicle(Emergency_Id int, User_Id int) {
	DB := dif.ConnectDB()
	Query := fmt.Sprintf("update dispatched_vehicles set user_id='%d' where emergency_id='%d'", User_Id, Emergency_Id)
	fmt.Println("amb1", Query)
	_ = DB.QueryRow(Query)

	Query = fmt.Sprintf("select lat,long,name,phone,type,updated_description from emergency where id='%d'", Emergency_Id)
	fmt.Println("amb2", Query)
	row := DB.QueryRow(Query)
	vehicle_notification := VehicleNotificationData{}
	_ = row.Scan(&vehicle_notification.Lat, &vehicle_notification.Long, &vehicle_notification.Name, &vehicle_notification.Phone, &vehicle_notification.Type, &vehicle_notification.Updated_Description)

	Query = fmt.Sprintf("select vehicle_id from dispatched_vehicles where emergency_id='%d'", Emergency_Id)
	fmt.Println("amb3", Query)
	rows, _ := DB.Query(Query)
	vehicle := Vehicle_Id{}

	for rows.Next() {
		_ = rows.Scan(&vehicle.Id)
		Query = fmt.Sprintf("update vehicle_data set status=false where id='%d'", vehicle.Id)
		fmt.Println("amb4", Query)
		_ = DB.QueryRow(Query)

		Query = fmt.Sprintf("select lat,long from vehicle_data where id='%d'", vehicle.Id)
		fmt.Println("amb5", Query)
		row = DB.QueryRow(Query)
		_ = row.Scan(&vehicle_notification.Vehicle_Lat, &vehicle_notification.Vehicle_Long)

		Query = fmt.Sprintf("select token from vehicle_token_data where vehicle_id='%d'", vehicle.Id)
		fmt.Println("amb6", Query)
		row = DB.QueryRow(Query)
		_ = row.Scan(&vehicle_notification.Token)

		go VehicleFCM(vehicle_notification)

	}

}
