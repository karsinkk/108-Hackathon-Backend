package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/karsinkk/108/controllers/admincontroller"
	"github.com/karsinkk/108/controllers/usercontroller"
	"github.com/karsinkk/108/controllers/vehiclecontroller"
)

func main() {
	router := mux.NewRouter()
	VehicleAppRouter := router.PathPrefix("/vehicle").Subrouter()
	UserAppRouter := router.PathPrefix("/user").Subrouter()
	AdminAppRouter := router.PathPrefix("/admin").Subrouter()

	VehicleAppRouter.HandleFunc("/login", vehiclecontroller.LoginVehicle).Methods("POST")
	VehicleAppRouter.HandleFunc("/update", vehiclecontroller.UpdateVehicle)
	VehicleAppRouter.HandleFunc("/notification", vehiclecontroller.UpdateVehicle)
	VehicleAppRouter.HandleFunc("/finish", vehiclecontroller.Finish)

	UserAppRouter.HandleFunc("/emergency", usercontroller.Emergency).Methods("POST")

	AdminAppRouter.HandleFunc("/register", admincontroller.Register).Methods("POST")
	AdminAppRouter.HandleFunc("/seen", admincontroller.ModifySeen)
	AdminAppRouter.HandleFunc("/addvehicle", admincontroller.AddVehicle).Methods("POST")
	AdminAppRouter.HandleFunc("/login", admincontroller.Login).Methods("POST")
	AdminAppRouter.HandleFunc("/notification", admincontroller.Notification)
	AdminAppRouter.HandleFunc("/emergency", admincontroller.DisplayEmergency)
	AdminAppRouter.HandleFunc("/emergencycount", admincontroller.CountEmergency)
	AdminAppRouter.HandleFunc("/status", admincontroller.Status).Methods("POST")
	AdminAppRouter.HandleFunc("/dismiss", admincontroller.DismissEmergency).Methods("POST")
	AdminAppRouter.HandleFunc("/dismissemergency", admincontroller.Dismiss).Methods("POST")
	AdminAppRouter.HandleFunc("/ambulance", admincontroller.DisplayAmbulance)
	AdminAppRouter.HandleFunc("/firepolice", admincontroller.DisplayFirePolice)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":4000")
}
