package admincontroller

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/karsinkk/108/dif"
	"github.com/karsinkk/108/helpers"
	"net/http"
	"time"
)

func SendNotifications(conn *websocket.Conn, base_id int) {

	DB := dif.ConnectDB()
	Query := fmt.Sprintf("select emergency.id eid,emergency.lat elat,emergency.long elong,emergency.phone ephone,emergency.name ename ,emergency.time etime ,emergency.status estatus,emergency.type etype ,emergency.description edescription,emergency.seen eseen ,emergency.updated_time eupdated ,d.vehicle_id vid,d.time_taken vtime ,v.district district,v.name vname,v.phone vphone,v.lat vlat,v.long vlong,v.driver driver,v.vehicle_no vno from emergency join dispatched_vehicles d on emergency.id=d.emergency_id join vehicle_data v on v.id=d.vehicle_id where emergency.status=TRUE and v.status=TRUE order by emergency.id")
	rows, _ := DB.Query(Query)
	defer rows.Close()
	data := helpers.Notification{}
	datas := make([]helpers.Notification, 0)
	for rows.Next() {

		if err := rows.Scan(&data.Eid, &data.ELat, &data.ELong, &data.Phone_1, &data.Name_1, &data.Time, &data.Status, &data.Type, &data.Description, &data.Seen, &data.Updated_time, &data.Vehicle_id, &data.Time_taken, &data.District, &data.Name_2, &data.Phone_2, &data.VLat, &data.VLong, &data.Driver, &data.Vehicle_no); err != nil {
			fmt.Println(err)
		}
		datas = append(datas, data)
	}
	conn.WriteJSON(datas)
	ticker1 := time.NewTicker(time.Millisecond * 30000)
	for _ = range ticker1.C {
		rows, _ = DB.Query(Query)
		defer rows.Close()
		data = helpers.Notification{}
		datas = make([]helpers.Notification, 0)
		for rows.Next() {

			if err := rows.Scan(&data.Eid, &data.ELat, &data.ELong, &data.Phone_1, &data.Name_1, &data.Time, &data.Status, &data.Type, &data.Description, &data.Seen, &data.Updated_time, &data.Vehicle_id, &data.Time_taken, &data.District, &data.Name_2, &data.Phone_2, &data.VLat, &data.VLong, &data.Driver, &data.Vehicle_no); err != nil {
				fmt.Println(err)
			}
			datas = append(datas, data)
		}
		conn.WriteJSON(datas)

	}
}
func Notification(res http.ResponseWriter, req *http.Request) {

	conn, err := helpers.Upgrader.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {

		var base_id helpers.AdminNotificationPostData
		err := conn.ReadJSON(&base_id)

		if err != nil {
			fmt.Println(err)
			conn.Close()
			return
		}

		// str := fmt.Sprintf("%+v", base_id)
		// fmt.Println(str)
		go SendNotifications(conn, base_id.Id)

	}()

}
