package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/karsinkk/108/controllers/admincontroller"
	"github.com/karsinkk/108/controllers/usercontroller"
	"github.com/karsinkk/108/controllers/vehiclecontroller"
	"github.com/karsinkk/108/ws"
	// "golang.org/x/net/websocket"
	// "log"
	// "net/http"
)

func main() {
	router := mux.NewRouter()
	VehicleAppRouter := router.PathPrefix("/vehicle").Subrouter()
	UserAppRouter := router.PathPrefix("/user").Subrouter()
	AdminAppRouter := router.PathPrefix("/admin").Subrouter()

	VehicleAppRouter.HandleFunc("/register", vehiclecontroller.RegisterVehicle).Methods("POST")
	VehicleAppRouter.HandleFunc("/update", vehiclecontroller.UpdateVehicle).Methods("POST")
	// AmbAppRouter.HandleFunc("/finish", AmbApp.Finish).Methods("POST")
	// AmbAppRouter.HandleFunc("/duty", AmbApp.Duty).Methods("POST")

	UserAppRouter.HandleFunc("/emergency", usercontroller.Emergency).Methods("POST")
	UserAppRouter.HandleFunc("/rating", usercontroller.SubmitRating).Methods("POST")

	AdminAppRouter.HandleFunc("/register", admincontroller.Register).Methods("POST")
	AdminAppRouter.HandleFunc("/login", admincontroller.Login).Methods("POST")
	AdminAppRouter.HandleFunc("/emergency/monthwise", admincontroller.ListEmergencyMonthWise).Methods("GET")
	AdminAppRouter.HandleFunc("/emergency", admincontroller.DisplayEmergency).Methods("GET")
	AdminAppRouter.HandleFunc("/ambulance/count", admincontroller.CountAmbulances).Methods("GET")
	AdminAppRouter.HandleFunc("/ambulance", admincontroller.ListAmbulances).Methods("GET")
	AdminAppRouter.HandleFunc("/notification", admincontroller.Notification).Methods("GET")
	AdminAppRouter.HandleFunc("/notif", ws.Notification)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":4000")
}
