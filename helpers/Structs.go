package helpers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/karsinkk/108-Hackathon-Backend/dif"
)

// GetUpgrader returns a WebSocket upgrader with configurable origin checking
func GetUpgrader() websocket.Upgrader {
	config := dif.GetConfig()
	allowedOrigins := strings.Split(config.AllowedOrigins, ",")

	return websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			if origin == "" {
				return true
			}

			// Allow all origins if configured with "*"
			if len(allowedOrigins) == 1 && allowedOrigins[0] == "*" {
				return true
			}

			// Check if origin is in allowed list
			for _, allowed := range allowedOrigins {
				if strings.TrimSpace(allowed) == origin {
					return true
				}
			}
			return false
		},
	}
}

// Upgrader is kept for backward compatibility
// Deprecated: Use GetUpgrader() instead for configurable CORS
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
	Username string `json:"username"`
	Password string `json:"password"`
}

type VehicleLoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Lat      string `json:"lat"`
	Long     string `json:"long"`
	Token    string `json:"token"`
}

type VehicleUpdateData struct {
	Id    int    `json:"Id"`
	Lat   string `json:"Lat"`
	Long  string `json:"Long"`
	Token string `json:"Token"`
}

type EmergencyUserData struct {
	Lat         string `json:"lat"`
	Long        string `json:"long"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Type        int    `json:"type"`
	Description int    `json:"description"`
	Number      int    `json:"number"`
	Token       string `json:"token"`
}

type EmergencyData struct {
	Id                  int       `json:"id"`
	Lat                 string    `json:"lat"`
	Long                string    `json:"long"`
	Phone               string    `json:"phone"`
	Name                string    `json:"name"`
	Status              bool      `json:"status"`
	Time                time.Time `json:"time"`
	Type                int       `json:"type"`
	Description         int       `json:"description"`
	Seen                bool      `json:"seen"`
	Updated_time        string    `json:"updated_time"`
	Updated_description string    `json:"updated_description"`
	Dismissed           bool      `json:"dismissed"`
}

type Vehicle_Id struct {
	Id int `json:"id"`
}

type AdminNotificationPostData struct {
	Id int `json:"Id"`
}

type VehicleWithinRadius struct {
	Id       string  `json:"id"`
	Lat      string  `json:"lat"`
	Long     string  `json:"long"`
	Distance float64 `json:"distance"`
}

type VehicleData struct {
	Id         int     `json:"id"`
	Lat        string  `json:"lat"`
	Long       string  `json:"long"`
	Phone      string  `json:"phone"`
	Driver     string  `json:"driver"`
	Vehicle_no string  `json:"vehicle_no"`
	Distance   float64 `json:"distance"`
	Time       int     `json:"time"`
}

type AdminRegisterData struct {
	Username string `json:"username"`
}

type HasuraSignupData struct {
	AuthToken   string   `json:"auth_token"`
	HasuraID    int      `json:"hasura_id"`
	HasuraRoles []string `json:"hasura_roles"`
}

type SignUpReturnData struct {
	Auth_token string `json:"auth_token"`
	Auth       string `json:"auth"`
	Id         int    `json:"id"`
}

type Emergencies struct {
	Id           int       `json:"id"`
	Hosp_id      int       `json:"hosp_id"`
	Ambulance_id int       `json:"ambulance_id"`
	Lat          string    `json:"lat"`
	Long         string    `json:"long"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Status       bool      `json:"status"`
	Time         time.Time `json:"time"`
}

type Ambulance_Count struct {
	OffDuty int `json:"off_duty"`
	OnDuty  int `json:"on_duty"`
}

type EmergencyMonthWise struct {
	Month string `json:"month"`
	Year  int    `json:"year"`
	Id    int    `json:"id"`
}

type AmbRegisterData struct {
	Id         int    `json:"id"`
	Vehicle_no string `json:"vehicle_no"`
}

type HasuraLoginData struct {
	HasuraRoles []string `json:"hasura_roles"`
	HasuraID    int      `json:"hasura_id"`
	AuthToken   string   `json:"auth_token"`
}

type LoginReturnData struct {
	Auth string `json:"auth"`
}

type Rating struct {
	Id     int `json:"id"`
	Rating int `json:"rating"`
}

type Notification struct {
	Eid          int       `json:"eid"`
	ELat         string    `json:"e_lat"`
	ELong        string    `json:"e_long"`
	Phone_1      string    `json:"phone_1"`
	Name_1       string    `json:"name_1"`
	Time         time.Time `json:"time"`
	Status       bool      `json:"status"`
	Type         int       `json:"type"`
	Description  int       `json:"description"`
	Seen         bool      `json:"seen"`
	Updated_time time.Time `json:"updated_time"`
	Vehicle_id   int       `json:"vehicle_id"`
	Time_taken   int       `json:"time_taken"`
	District     string    `json:"district"`
	Name_2       string    `json:"name_2"`
	Phone_2      string    `json:"phone_2"`
	VLat         string    `json:"v_lat"`
	VLong        string    `json:"v_long"`
	Driver       string    `json:"driver"`
	Vehicle_no   string    `json:"vehicle_no"`
}

type StatusData struct {
	Dispatched          bool   `json:"dispatched"`
	Emergency_Id        int    `json:"emergency_id"`
	User_Id             int    `json:"user_id"`
	Updated_Description string `json:"updated_description"`
}

type VehicleNotificationData struct {
	Lat                 string `json:"lat"`
	Long                string `json:"long"`
	Phone               string `json:"phone"`
	Name                string `json:"name"`
	Type                int    `json:"type"`
	Updated_Description string `json:"updated_description"`
	Vehicle_Lat         string `json:"vehicle_lat"`
	Vehicle_Long        string `json:"vehicle_long"`
	Token               string `json:"token"`
}

type DismissData struct {
	Vehicle_Id   int `json:"vehicle_id"`
	Emergency_Id int `json:"emergency_id"`
}

type UserNotificationData struct {
	Lat        string `json:"lat"`
	Long       string `json:"long"`
	Time       int    `json:"time"`
	Vehicle_No string `json:"vehicle_no"`
	Phone      string `json:"phone"`
	Name       string `json:"name"`
	Token      string `json:"token"`
}

type VehicleAddData struct {
	Vehicle_no string `json:"vehicle_no"`
	Driver     string `json:"driver"`
	Phone      string `json:"phone"`
}

type SeenData struct {
	Id []int `json:"id"`
}

type Vehicle struct {
	Id         int    `json:"id"`
	District   string `json:"district"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Lat        string `json:"lat"`
	Long       string `json:"long"`
	Driver     string `json:"driver"`
	Vehicle_no string `json:"vehicle_no"`
	Username   string `json:"username"`
	Status     bool   `json:"status"`
	Type       int    `json:"type"`
}

type DismissEmergencyData struct {
	Emergency_Id     int    `json:"emergency_id"`
	Dismissed_Reason string `json:"dismissed_reason"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Database  string `json:"database"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
