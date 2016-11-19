package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type apiPostData struct {
	Lat   string
	Long  string
	Tag   string
	Etype string
}

type test_struct struct {
	Test string
}

func main() {

	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./108-Hackathon-Website/")))).Methods("GET")
	router.HandleFunc("/api", handleAppPost).Methods("POST")
	http.Handle("/", router)
	http.ListenAndServe(":80", router)
}

func handleAppGet(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(200)
	fmt.Fprintf(res, "GT Thevdiya")
}
func handleAppPost(res http.ResponseWriter, req *http.Request) {
	var t apiPostData
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	fmt.Println(t)
	log.Println(t.Lat)
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(200)
	fmt.Fprintf(res, t.Lat)
}
