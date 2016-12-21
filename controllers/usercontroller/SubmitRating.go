package usercontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
	"net/http"
)

func SubmitRating(res http.ResponseWriter, req *http.Request) {
	DB := dif.ConnectDB()
	var u helpers.Rating
	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	Query := fmt.Sprintf("update hospital set rating=(((rating*rating_count)+(%d))/(rating_count+1)),rating_count=(rating_count+1) where id=%d", u.Rating, u.Id)
	fmt.Println(Query)
	_, _ = DB.Query(Query)

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("108", "An NP-Incomplete Project")
	res.WriteHeader(200)

}
