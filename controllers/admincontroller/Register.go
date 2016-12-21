package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"net/http"
)

func Register(res http.ResponseWriter, req *http.Request) {
	var r helpers.AdminRegisterData
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	str := fmt.Sprintf("%+v \n", r)
	fmt.Print(str)

	var data helpers.SignUpReturnData
	// var header map[string][]string
	res.Header().Set("Chiron", "An NP-Incomplete Project")
	data.Auth_token = helpers.RegisterUser(r)

	res.WriteHeader(200)
	// for i, j := range header {
	// 	for _, l := range j {
	// 		fmt.Println(i, l)
	// 		res.Header().Set(i, l)
	// 	}
	// }
	res.Header().Set("Access-Control-Allow-Origin", "*")
	str = fmt.Sprintf("%+v \n", data)
	fmt.Print(str)
	json.NewEncoder(res).Encode(data)

}
