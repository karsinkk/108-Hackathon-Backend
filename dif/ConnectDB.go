package dif

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	PsqlInfo := fmt.Sprintf("host=localhost port=5432 user=postgres dbname=dev sslmode=disable")
	DB, _ := sql.Open("postgres", PsqlInfo)
	return DB
}
