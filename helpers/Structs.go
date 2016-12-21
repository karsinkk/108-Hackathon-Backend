package helpers

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Duration struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type Distance struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type DistanceMatrixElement struct {
	Duration Duration `json:"duration"`
	Distance Distance `json:"distance"`
	Status   string   `json:"status"`
}

type DistanceMatrixElementsRow struct {
	Elements []DistanceMatrixElement `json:"elements"`
}

type DistanceMatrixResponse struct {
	OriginAddresses      []string                    `json:"origin_addresses"`
	DestinationAddresses []string                    `json:"destination_addresses"`
	Rows                 []DistanceMatrixElementsRow `json:"rows"`
}

type MIN struct {
	Key   int
	Value int
}

type LoginData struct {
	Username string
	Password string
}
type VehicleLoginData struct {
	Username string
	Password string
	Lat      string
	Long     string
	Token    string
}

type VehicleUpdateData struct {
	Id    int    `json:"Id"`
	Lat   string `json:"Lat"`
	Long  string `json:"Long"`
	Token string `json:"Token"`
}

type EmergencyUserData struct {
	Lat         string
	Long        string
	Name        string
	Phone       string
	Type        int
	Description int
	Number      int
	Token       string
}
type EmergencyData struct {
	Id                  int
	Lat                 string
	Long                string
	Phone               string
	Name                string
	Status              bool
	Time                time.Time
	Type                int
	Description         int
	Seen                bool
	Updated_time        string
	Updated_description string
	Dismissed           bool
}

type Vehicle_Id struct {
	Id int
}

type AdminNotificationPostData struct {
	Id int `json:"Id"`
}

type VehicleWithinRadius struct {
	Id       string
	Lat      string
	Long     string
	Distance float64
}

type VehicleData struct {
	Id         int
	Lat        string
	Long       string
	Phone      string
	Driver     string
	Vehicle_no string
	Distance   float64
	Time       int
}

type AdminRegisterData struct {
	Username string
}

type HasuraSignupData struct {
	AuthToken   string   `json:"auth_token"`
	HasuraID    int      `json:"hasura_id"`
	HasuraRoles []string `json:"hasura_roles"`
}

type SignUpReturnData struct {
	Auth_token string
	Auth       string
	Id         int
}

type Emergencies struct {
	Id           int
	Hosp_id      int
	Ambulance_id int
	Lat          string
	Long         string
	Name         string
	Phone        string
	Status       bool
	Time         time.Time
}

type Ambulance_Count struct {
	OffDuty int
	OnDuty  int
}

type EmergencyMonthWise struct {
	Month string
	Year  int
	Id    int
}

type AmbRegisterData struct {
	Id         int
	Vehicle_no string
}

type HasuraLoginData struct {
	HasuraRoles []string `json:"hasura_roles"`
	HasuraID    int      `json:"hasura_id"`
	AuthToken   string   `json:"auth_token"`
}

type LoginReturnData struct {
	Auth string
}

type Rating struct {
	Id     int
	Rating int
}

type Notification struct {
	Eid          int
	ELat         string
	ELong        string
	Phone_1      string
	Name_1       string
	Time         time.Time
	Status       bool
	Type         int
	Description  int
	Seen         bool
	Updated_time time.Time
	Vehicle_id   int
	Time_taken   int
	District     string
	Name_2       string
	Phone_2      string
	VLat         string
	VLong        string
	Driver       string
	Vehicle_no   string
}

type StatusData struct {
	Dispatched          bool
	Emergency_Id        int
	User_Id             int
	Updated_Description string
}

type VehicleNotificationData struct {
	Lat                 string
	Long                string
	Phone               string
	Name                string
	Type                int
	Updated_Description string
	Vehicle_Lat         string
	Vehicle_Long        string
	Token               string
}
type DismissData struct {
	Vehicle_Id   int
	Emergency_Id int
}

type UserNotificationData struct {
	Lat        string
	Long       string
	Time       int
	Vehicle_No string
	Phone      string
	Name       string
	Token      string
}

type VehicleAddData struct {
	Vehicle_no string
	Driver     string
	Phone      string
}

type SeenData struct {
	Id []int
}

type Vehicle struct {
	Id         int
	District   string
	Name       string
	Phone      string
	Lat        string
	Long       string
	Driver     string
	Vehicle_no string
	Username   string
	Status     bool
	Type       int
}

type DismissEmergencyData struct {
	Emergency_Id     int
	Dismissed_Reason string
}
