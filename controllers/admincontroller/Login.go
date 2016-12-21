package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"net/http"
	"strconv"
	"strings"
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
	res.Header().Set("108", "An NP-Incomplete Project")
	data.Auth = helpers.LoginUser(r)

	res.WriteHeader(200)
	res.Header().Set("Access-Control-Allow-Origin", "*")
	str = fmt.Sprintf("%+v \n", data)
	fmt.Print(str)

	lol, _ := strconv.ParseInt(strings.Split(r.Username, "adminuser")[1], 10, 64)
	fmt.Fprint(res, lol)
	// json.NewEncoder(res).Encode(data)

}
