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

type ApiUserPostData struct {
	Lat   string
	Long  string
	Tag   string
	Etype string
}

type ApiAmbPostData struct {
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
	EType    string
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		forward := req.Header.Get("X-Forwarded-For")
		log.Printf("%s %s %s \n\n", forward, req.Method, req.URL)
		handler.ServeHTTP(res, req)
	})
}

func ReadLine(lineNum int) (line string) {
	r, _ := os.Open("BaseData.csv")
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

var EType = map[string]string{"medical": "Ambulance", "fire": "Fire Engine", "police": "Police Vehicle"}

func main() {

	ReadConf()

	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./108-Hackathon-Website/")))).Methods("GET")
	router.HandleFunc("/api/user", handleAppUserPost).Methods("POST")
	router.HandleFunc("/api/amb", handleAppAmbPost).Methods("POST")
	http.Handle("/", router)
	http.ListenAndServe(":8080", Log(http.DefaultServeMux))
}

func handleAppAmbPost(res http.ResponseWriter, req *http.Request) {
	var t ApiAmbPostData
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	str := fmt.Sprintf("%+v \n", t)
	fmt.Print(str)

	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(200)
	fmt.Fprintf(res, "Thanks!!\n")
}
func handleAppUserPost(res http.ResponseWriter, req *http.Request) {
	var u ApiUserPostData
	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	str := fmt.Sprintf("%+v \n", u)
	fmt.Print(str)

	ID, Time := NearestBase(u.Lat, u.Long)
	BaseData := ReadLine(ID + 1)
	fmt.Println(BaseData, "Time:", Time)
	fmt.Printf("ID:%d Emergency Services can reach you in %d mins\n\n", ID, Time/60)

	BD := strings.Split(BaseData, ",")
	BData := BaseLocation{District: BD[0], Locality: BD[1], Lat: BD[2], Long: BD[3], Time: Time / 60, EType: EType[u.Etype]}

	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(BData)
}
