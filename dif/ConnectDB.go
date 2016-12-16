package dif

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	Conf := ReadConf()
	PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"dbname=%s sslmode=disable", Conf.Host, Conf.Port, Conf.Username, Conf.DBname)
	DB, _ := sql.Open("postgres", PsqlInfo)
	return DB
}
