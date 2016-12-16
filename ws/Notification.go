package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/karsinkk/108/dif"
	"net/http"
	"time"
)

type Data struct {
	Lat  string
	Long string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Notification(res http.ResponseWriter, req *http.Request) {
	var err error
	DB := dif.ConnectDB()
	Query := fmt.Sprintf("select lat,long from vehicles")
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		ticker1 := time.NewTicker(time.Millisecond * 1000)
		for _ = range ticker1.C {
			rows, _ := DB.Query(Query)
			defer rows.Close()
			data := Data{}
			datas := make([]Data, 0)
			for rows.Next() {

				if err := rows.Scan(&data.Lat, &data.Long); err != nil {
					fmt.Println(err)
				}
				datas = append(datas, data)
			}
			conn.WriteJSON(datas)

		}
		fmt.Println("In stop")
	}()

	go func() {
		for {
			var reply []byte
			msgType, reply, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				conn.Close()
				return
			}

			fmt.Println("Received back from client: " + string(reply[:]))

			msg := "Received:  " + string(reply[:])
			fmt.Println("Sending to client: " + msg)
			for i := 0; i < 1; i++ {
				if err = conn.WriteMessage(msgType, reply); err != nil {
					fmt.Println("Can't send")
					break
				}
			}
		}
	}()
}
