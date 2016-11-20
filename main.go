package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var f, _ = os.OpenFile("AppPostData.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)

type apiPostData struct {
	Lat   string
	Long  string
	Tag   string
	Etype string
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		forward := req.Header.Get("X-Forwarded-For")
		log.Printf("%s %s %s \n", forward, req.Method, req.URL)
		handler.ServeHTTP(res, req)
	})
}

func main() {

	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./108-Hackathon-Website/")))).Methods("GET")
	router.HandleFunc("/api", handleAppPost).Methods("POST")
	http.Handle("/", router)
	http.ListenAndServe(":8080", Log(http.DefaultServeMux))
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
	str := fmt.Sprintf("%+v \n", t)
	f.WriteString(str)
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(200)
	fmt.Fprintf(res, str)
}
