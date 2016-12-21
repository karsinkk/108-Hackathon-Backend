package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	PsqlInfo := fmt.Sprintf("host=localhost port=5432 user=postgres dbname=dev sslmode=disable")
	DB, _ := sql.Open("postgres", PsqlInfo)
	err := DB.Ping()

	if err != nil {
		fmt.Println("Fucked up")
	}
	Query := fmt.Sprintf("select id from emergency")
	rows, _ := DB.Query(Query)
	var id int
	ids := make([]int, 0)
	for rows.Next() {
		_ := rows.Scan(&id)
		ids = append(ids, id)
	}

	fmt.Println(ids)
}
