package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
)

func RegisterUser(data AdminRegisterData) string {

	values := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, data.Username, data.Username)
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

	var r HasuraSignupData

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		fmt.Println(err)
	}

	auth := r.AuthToken[0:]

	return auth
}
