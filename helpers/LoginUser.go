package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginUser(data LoginData) string {
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

	fmt.Println(resp.Header)
	var r HasuraLoginData

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		fmt.Println(err)
	}

	auth := r.AuthToken[0:]

	return auth
}
