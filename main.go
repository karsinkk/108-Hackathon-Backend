package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
)

var f, _ = os.OpenFile("AppPostData.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)

type apiPostData struct {
	Lat   string
	Long  string
	Tag   string
	Etype string
}

type BaseLocation struct {
	District string
	Locality string
	Lat      string
	Long     string
	Time     int
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		forward := req.Header.Get("X-Forwarded-For")
		log.Printf("%s %s %s \n\n", forward, req.Method, req.URL)
		handler.ServeHTTP(res, req)
	})
}

func ReadLine(lineNum int) (line string) {
	r, _ := os.Open("Base Location LAT & LONG.csv")
	lastLine := 0
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			return sc.Text()
		}
	}
	defer r.Close()
	return line
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
	fmt.Print(str)

	ID, Time := NearestBase(t.Lat, t.Long)
	BaseData := ReadLine(ID)
	fmt.Println(BaseData, "Time:", Time)
	fmt.Printf("ID:%d Emergency Services can reach you in %d mins\n\n", ID, Time/60)

	BD := strings.Split(BaseData, ",")
	BData := BaseLocation{District: BD[0], Locality: BD[1], Lat: BD[2], Long: BD[3], Time: Time / 60}

	f.WriteString(str)

	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(BData)
}
