package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/Chiron-Backend/helpers"
	"net/http"
)

func Login(res http.ResponseWriter, req *http.Request) {
	var r helpers.AdminLoginData
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	str := fmt.Sprintf("%+v \n", r)
	fmt.Print(str)

	var data helpers.LoginReturnData
	res.Header().Set("Chiron", "An NP-Incomplete Project")
	data.Auth = helpers.LoginUser(r)

	res.WriteHeader(200)
	res.Header().Set("Access-Control-Allow-Origin", "*")
	str = fmt.Sprintf("%+v \n", data)
	fmt.Print(str)
	json.NewEncoder(res).Encode(data)

}
