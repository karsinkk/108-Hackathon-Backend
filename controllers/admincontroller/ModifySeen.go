package admincontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108/dif"
	"github.com/karsinkk/108/helpers"
	"net/http"
)

func ModifySeen(res http.ResponseWriter, req *http.Request) {
	DB := dif.ConnectDB()
	var u helpers.SeenData
	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	str := fmt.Sprintf("%+v \n", u)
	fmt.Print(str)
	if len(u.Id) > 0 {
		s := fmt.Sprintf("%d", u.Id[0])
		for _, v := range u.Id {
			s += fmt.Sprintf(",%d", v)
		}
		Query := fmt.Sprintf("update emergency set seen=true where id in(%s)", s)
		fmt.Println(Query)
		_ = DB.QueryRow(Query)
	}
	res.Header().Set("108", "An NP-Incomplete Project")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(200)
}
