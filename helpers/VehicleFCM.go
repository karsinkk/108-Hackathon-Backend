package helpers

import (
	"bytes"
	"fmt"
	"github.com/karsinkk/108/dif"
	"net/http"
)

func VehicleFCM(vehicle_notification VehicleNotificationData) {
	Conf := dif.ReadConf()
	FCM_Key := "key=" + Conf.Vehicle_FCM_Key
	values := fmt.Sprintf(`{"data":{"Lat":"%s","Long":"%s","Vehicle_Lat":"%s","Vehicle_Long":"%s","Name":"%s","Phone":"%s","Updated_Description":"%s"},"to":"%s"}`, vehicle_notification.Lat, vehicle_notification.Long, vehicle_notification.Vehicle_Lat, vehicle_notification.Vehicle_Long, vehicle_notification.Name, vehicle_notification.Phone, vehicle_notification.Updated_Description, vehicle_notification.Token)
	fmt.Println("amb7", values)
	fmt.Println(FCM_Key)
	url := "https://fcm.googleapis.com/fcm/send"

	var jsonStr = []byte(values)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", FCM_Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
