package helpers

import (
	"bytes"
	"fmt"
	"github.com/karsinkk/108/dif"
	"net/http"
)

func UserFCM(user_notification UserNotificationData) {
	Conf := dif.ReadConf()
	FCM_Key := "key=" + Conf.User_FCM_Key
	values := fmt.Sprintf(`{"data":{"Lat":"%s","Long":"%s","Time":"%d","Vehicle_no":"%s","Phone":"%s","Name":"%s"},"to":"%s"}`, user_notification.Lat, user_notification.Long, user_notification.Time, user_notification.Vehicle_No, user_notification.Phone, user_notification.Name, user_notification.Token)
	fmt.Println(values)
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
