package admincontroller

import (
	"fmt"
	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"net/http"
	"time"
)

type Count struct {
	Id int
}

func CountEmergency(res http.ResponseWriter, req *http.Request) {

	conn, err := helpers.Upgrader.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	DB := dif.ConnectDB()
	Query := fmt.Sprintf("select count(id) from emergency where seen=false")
	count_data := Count{}
	row := DB.QueryRow(Query)
	_ = row.Scan(&count_data.Id)
	conn.WriteJSON(count_data)
	go func() {
		ticker1 := time.NewTicker(time.Millisecond * 30000)
		for _ = range ticker1.C {
			count_data = Count{}
			row = DB.QueryRow(Query)
			_ = row.Scan(&count_data.Id)
			conn.WriteJSON(count_data)

		}
	}()
}
