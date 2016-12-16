package helpers

import (
	"time"
)

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

type AmbRegisterToken struct {
	Auth_Token string
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

type Ambulance struct {
	Hosp_id    int
	Id         int
	Lat        string
	Long       string
	Phone      string
	Time       time.Time
	Status     bool
	Driver     string
	Vehicle_no string
	Distance   float64
}

type AmbulanceWithinRadius struct {
	Id       string
	Lat      string
	Long     string
	Distance float64
}

type VehicleData struct {
	Base_Id    int
	Id         string
	Lat        string
	Long       string
	Phone      string
	Driver     string
	Vehicle_no string
	Distance   float64
	Time       int
}

type EmergencyData struct {
	Hosp_id int
	Amb_id  int
	Lat     string
	Long    string
	Phone   string
	Name    string
}

type AdminRegisterData struct {
	Name     string
	Username string
	Phone    string
	Password string
	Address  string
}

type SignupData struct {
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

type AmbUpdateData struct {
	Id     int
	Lat    string
	Long   string
	Phone  string
	Status bool
	Driver string
}

type AdminLoginData struct {
	Username string
	Password string
}

type LoginData struct {
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
	Amb_id     int
	Driver     string
	Vehicle_no string
	Name       string
	Phone      string
}
