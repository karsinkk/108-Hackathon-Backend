package admincontroller

import (
	"fmt"
	"github.com/karsinkk/108/dif"
	"github.com/karsinkk/108/helpers"
	"net/http"
	"time"
)

func DisplayFirePolice(res http.ResponseWriter, req *http.Request) {

	conn, err := helpers.Upgrader.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	DB := dif.ConnectDB()
	Query := fmt.Sprintf("select * from emergency where type!=1 and status=true")
	data := helpers.EmergencyData{}
	datas := make([]helpers.EmergencyData, 0)
	rows, _ := DB.Query(Query)

	for rows.Next() {

		if err := rows.Scan(&data.Id, &data.Lat, &data.Long, &data.Phone, &data.Name, &data.Status, &data.Time, &data.Type, &data.Description, &data.Seen, &data.Updated_time, &data.Updated_description, &data.Dismissed); err != nil {
			fmt.Println(err)
		}
		datas = append(datas, data)
	}
	conn.WriteJSON(datas)
	go func() {
		ticker1 := time.NewTicker(time.Millisecond * 30000)
		for _ = range ticker1.C {
			data = helpers.EmergencyData{}
			datas = make([]helpers.EmergencyData, 0)
			rows, _ = DB.Query(Query)

			for rows.Next() {

				if err = rows.Scan(&data.Id, &data.Lat, &data.Long, &data.Phone, &data.Name, &data.Status, &data.Time, &data.Type, &data.Description, &data.Seen, &data.Updated_time, &data.Updated_description, &data.Dismissed); err != nil {
					fmt.Println(err)
				}
				datas = append(datas, data)
			}
			conn.WriteJSON(datas)
		}
	}()
}
