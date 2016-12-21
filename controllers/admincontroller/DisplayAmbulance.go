package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"net/http"
)

func DisplayAmbulance(res http.ResponseWriter, req *http.Request) {

	DB := dif.ConnectDB()
	Query := fmt.Sprintf("select * from vehicle_data where type=1")
	// conn, err := helpers.Upgrader.Upgrade(res, req, nil)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	rows, _ := DB.Query(Query)
	defer rows.Close()
	data := helpers.Vehicle{}
	datas := make([]helpers.Vehicle, 0)
	for rows.Next() {

		if err := rows.Scan(&data.Id, &data.District, &data.Name, &data.Phone, &data.Lat, &data.Long, &data.Driver, &data.Vehicle_no, &data.Username, &data.Status, &data.Type); err != nil {
			fmt.Println(err)
		}
		datas = append(datas, data)
	} //
	// conn.WriteJSON(datas)
	// go func() {
	// 	ticker1 := time.NewTicker(time.Millisecond * 30000)
	// 	for _ = range ticker1.C {
	// 		rows, _ = DB.Query(Query)
	// 		defer rows.Close()
	// 		data = helpers.Vehicle{}
	// 		datas = make([]helpers.Vehicle, 0)
	// 		for rows.Next() {
	//
	// 			if err = rows.Scan(&data.Id, &data.District, &data.Name, &data.Phone, &data.Lat, &data.Long, &data.Driver, &data.Vehicle_no, &data.Username, &data.Status, &data.Type); err != nil {
	// 				fmt.Println(err)
	// 			}
	// 			datas = append(datas, data)
	// 		}
	// 		conn.WriteJSON(datas)
	//
	// 	}
	// }()
	res.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(res).Encode(datas)

}
