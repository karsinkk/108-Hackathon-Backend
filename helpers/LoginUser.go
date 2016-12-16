package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108/dif"
	"net/http"
)

func LoginUser(data AdminLoginData) string {
	DB := dif.ConnectDB()
	values := fmt.Sprintf(`{"username": "%s","password": "%s"}`, data.Username, data.Password)
	url := "https://auth.archon40.hasura-app.io/login"

	var jsonStr = []byte(values)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var r LoginData

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		fmt.Println(err)
	}

	auth := r.AuthToken[0:]
	Query := fmt.Sprintf("update hospital set auth='%s' where email='%s'", auth, data.Username)
	fmt.Println(Query)
	_ = DB.QueryRow(Query)

	return auth
}
