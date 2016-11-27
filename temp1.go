package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Duration struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type Distance struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type DistanceMatrixElement struct {
	Duration Duration `json:"duration"`
	Distance Distance `json:"distance"`
	Status   string   `json:"status"`
}

type DistanceMatrixElementsRow struct {
	Elements []DistanceMatrixElement `json:"elements"`
}

type DistanceMatrixResponse struct {
	OriginAddresses      []string                    `json:"origin_addresses"`
	DestinationAddresses []string                    `json:"destination_addresses"`
	Rows                 []DistanceMatrixElementsRow `json:"rows"`
}

var BaseLocationDataEncoded = []string{"cbnnAuxshNqrPwcFxgBmnAw[hhJjcUql@{}Np~A}nJmmNvaV|_KihI}jIzqAbvDwvF{dF`xAxlDrcMxLgfLnYaoKgj@zbH|{L`xCixPzzTxiFij\\haA~bCxhA", "yq`oAitshNvuN~uA}}EwwK~v@vdA`{@`vGk|CmpChrIoqEajA|bH`a@gq@bxShcOzbNfdQit\\o~EnuiAjy`Ac`b@q`J{`h@uhjA|xWv}qBteLibMwrCuqCyCPwfJg|{@", "qs{mAca_hNznCnyy@bek@cw^p`l@kgDyodAuwBvkA}_e@t|T|{m@t{_@prPsnkA{uaAeW~{WweIpf^~kcBk{^{gkAjcnAzhk@}qdAklk@lmdAwwTgfsAfvbBz{y@qchBil}@dldAuqDcwRrirA", "kbklAycdgNqoXjoi@jcxA_{m@m`mBwad@`yFuqBsmo@haPjhu@saS~lO|yj@aue@br_AnloAm|l@aymAe`bAyniAvcIr_g@l|v@bjIciTudR`fdBy{AiedCji[t_SqhYrfzAfeLc_nBp{Mfi_A", "e}fpAilpgNqdIgsQty~@qcEctWchEylp@gyTrufAfjg@upl@dtxBn{YipoAhhBiwAgvMyer@oph@tlGjbt@yoPqtJnoyAxjNmlt@mpHn`oAnfLwagBqj`@mcTfza@hcmAuwl@}gFlhVikqAryQpjj@ubAvoc@vvKszp@"}

var API_KEY string = "AIzaSyCy-cCGCwE2qhtthSxAXMY71Z9Di7r2t2Y"
var BaseUrl string = "https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins="

var DurationData = make(map[int]int)

type MIN struct {
	Key   int
	Value int
}

func main() {
	var count int = 0
	var Min MIN
	Min.Value = 1<<31 - 1
	for _, Base := range BaseLocationDataEncoded {
		url := fmt.Sprintf("%s12.840639,80.170417&destinations=enc:%s:&key=%s", BaseUrl, Base, API_KEY)

		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var data DistanceMatrixResponse
		json.NewDecoder(resp.Body).Decode(&data)
		for _, i := range data.Rows {
			for _, j := range i.Elements {
				DValue := j.Duration.Value
				if DValue < Min.Value {
					Min.Key = count
					Min.Value = DValue
				}
				count++
			}
		}
		fmt.Println(Min, count)
	}
}
