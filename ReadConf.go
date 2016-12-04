package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Configuration struct {
	API_KEY                 string
	BaseLocationDataEncoded []string
}

var Conf = Configuration{}

func ReadConf() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(reflect.TypeOf(Conf.BaseLocationDataEncoded))
}
