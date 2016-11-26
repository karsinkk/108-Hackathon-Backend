package main

import (
	"fmt"
	"ioutil"
	"net/http"
)

var BaseLocationData = []string{"cbnnAuxshNqrPwcFxgBmnAw[hhJjcUql@{}Np~A}nJmmNvaV|_KihI}jIzqAbvDwvF{dF`xAxlDrcMxLgfLnYaoKgj@zbH|{L`xCixPzzTxiFij\\haA~bCxhA", "yq`oAitshNvuN~uA}}EwwK~v@vdA`{@`vGk|CmpChrIoqEajA|bH`a@gq@bxShcOzbNfdQit\\o~EnuiAjy`Ac`b@q`J{`h@uhjA|xWv}qBteLibMwrCuqCyCPwfJg|{@", "qs{mAca_hNznCnyy@bek@cw^p`l@kgDyodAuwBvkA}_e@t|T|{m@t{_@prPsnkA{uaAeW~{WweIpf^~kcBk{^{gkAjcnAzhk@}qdAklk@lmdAwwTgfsAfvbBz{y@qchBil}@dldAuqDcwRrirA", "kbklAycdgNqoXjoi@jcxA_{m@m`mBwad@`yFuqBsmo@haPjhu@saS~lO|yj@aue@br_AnloAm|l@aymAe`bAyniAvcIr_g@l|v@bjIciTudR`fdBy{AiedCji[t_SqhYrfzAfeLc_nBp{Mfi_A", "e}fpAilpgNqdIgsQty~@qcEctWchEylp@gyTrufAfjg@upl@dtxBn{YipoAhhBiwAgvMyer@oph@tlGjbt@yoPqtJnoyAxjNmlt@mpHn`oAnfLwagBqj`@mcTfza@hcmAuwl@}gFlhVikqAryQpjj@ubAvoc@vvKszp@"}

var API_KEY = "AIzaSyCy-cCGCwE2qhtthSxAXMY71Z9Di7r2t2Y"

func main() {
	BaseUrl := "https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins="
	url := fmt.Sprintf("%s12.840639,80.170417&destinations=enc:%s:&key=%s", BaseUrl, BaseLocationData[0], API_KEY)

	req, err := http.NewRequest("GET", url)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
