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

type Movie struct {
	Title  string `json:"title`
	Rating string `json:"rating"`
	Year   string `json:"year"`
}

func main() {

	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./108-Hackathon-Website/")))).Methods("GET")
	router.HandleFunc("/api", handleAppPost).Methods("POST")
	router.HandleFunc("/movies", handleMovies).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
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

func handleMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var Movies = map[string]*Movie{
		"00001": &Movie{Title: "Batman Begins", Rating: "9.8", Year: "2005"},
		"00002": &Movie{Title: "The Dark Knight", Rating: "9.9", Year: "2010"},
		"00003": &Movie{Title: "The Dark Knight Rises", Rating: "10.0", Year: "2015"},
	}
	json.NewEncoder(res).Encode(Movies)
}
