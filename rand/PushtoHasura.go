package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type BaseData struct {
	District string
	Locality string
	Lat      string
	Long     string
}
type Arg struct {
	table     string
	objects   []BaseData
	returning []string
}
type HasuraQuery struct {
	Type string
	Args Arg
}

func main() {

	f, _ := os.Open("Base Location LAT & LONG.csv")

	defer f.Close()

	csvr := csv.NewReader(f)

	var data = make(BaseData{}, 0)
	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		data = BaseData{District: row[0], Locality: row[1], Lat: row[2], Long: row[3]}
		fmt.Println(data)

	}

}
