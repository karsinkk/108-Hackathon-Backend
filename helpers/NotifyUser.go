package helpers

import (
	"fmt"
	"github.com/karsinkk/108/dif"
)

func NotifyUser(Emergency_Id int) {
	DB := dif.ConnectDB()

	Query := fmt.Sprintf("select vehicle_id,time_taken from dispatched_vehicles where emergency_id='%d' order by time_taken limit 1", Emergency_Id)
	fmt.Println(Query)
	row := DB.QueryRow(Query)
	var vehicle_id int
	user_notification := UserNotificationData{}
	_ = row.Scan(&vehicle_id, &user_notification.Time)
	user_notification.Time = user_notification.Time / 60
	Query = fmt.Sprintf("select lat,long,vehicle_no,phone,driver from vehicle_data where id='%d'", vehicle_id)
	fmt.Println(Query)
	row = DB.QueryRow(Query)
	_ = row.Scan(&user_notification.Lat, &user_notification.Long, &user_notification.Vehicle_No, &user_notification.Phone, &user_notification.Name)

	Query = fmt.Sprintf("select token from emergency_token_data where emergency_id='%d'", Emergency_Id)
	fmt.Println(Query)
	row = DB.QueryRow(Query)
	_ = row.Scan(&user_notification.Token)

	UserFCM(user_notification)

}
