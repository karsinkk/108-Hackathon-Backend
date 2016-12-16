package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108/dif"
	"net/http"
)

func RegisterUser(data AdminRegisterData) (string, string, int, http.Header) {
	DB := dif.ConnectDB()
	values := fmt.Sprintf(`{"username": "%s", "mobile": "%s", "password": "%s"}`, data.Username, data.Phone, data.Password)
	url := "https://auth.archon40.hasura-app.io/signup"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(values)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var r SignupData

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		fmt.Println(err)
	}

	auth_token := r.AuthToken[0:6]
	auth := r.AuthToken[0:]
	Query := fmt.Sprintf("insert into hospital(name,address,phone,email,auth_token,auth) values('%s','%s','%s','%s','%s','%s')", data.Name, data.Address, data.Phone, data.Username, auth_token, auth)
	_ = DB.QueryRow(Query)

	Query = fmt.Sprintf("select id from hospital where auth_token='%s'", auth_token)
	var id int
	_ = DB.QueryRow(Query).Scan(&id)
	fmt.Println(resp.Header)

	return auth_token, auth, id, resp.Header
}
