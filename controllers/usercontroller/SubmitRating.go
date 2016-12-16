package usercontroller

import (
	"encoding/json"
	"fmt"
	"github.com/karsinkk/Chiron-Backend/helpers"
	"net/http"
)

func SubmitRating(res http.ResponseWriter, req *http.Request) {
	var u helpers.Rating
	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	Query := fmt.Sprintf("update hospital set rating=(((rating*rating_count)+(%d))/(rating_count+1)),rating_count=(rating_count+1) where id=%d", u.Rating, u.Id)
	fmt.Println(Query)
	_, _ = helpers.DB.Query(Query)

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Chiron", "An NP-Incomplete Project")
	res.WriteHeader(200)

}
