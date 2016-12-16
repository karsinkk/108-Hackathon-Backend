package main

import (
	"fmt"
	// "github.com/karsinkk/108/helpers"
	"database/sql"
	_ "github.com/lib/pq"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

type Data struct {
	Lat  string
	Long string
}

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"dbname=%s sslmode=disable",
	host, port, username, dbname)
var DB, _ = sql.Open("postgres", psqlInfo)

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = ""
	dbname   = "dev"
)

func Echo(ws *websocket.Conn) {
	var err error

	Query := fmt.Sprintf("select lat,long from vehicles")
	fmt.Println(Query)
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

	websocket.JSON.Send(ws, datas)
	//
	// ticker1 := time.NewTicker(time.Millisecond * 1000)
	// // time.Sleep(time.Millisecond * 2400)
	// go func() {
	// 	for i := range ticker1.C {
	// 		message := fmt.Sprintf("Timer Message No. : %d", i)
	// 		websocket.Message.Send(ws, message)
	// 	}
	// }()

	// time.Sleep(time.Millisecond * 6000)
	// ticker1.Stop()

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)
		for i := 0; i < 1; i++ {
			if err = websocket.Message.Send(ws, msg); err != nil {
				fmt.Println("Can't send")
				break
			}
		}
	}
}

func main() {
	http.Handle("/notif", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
