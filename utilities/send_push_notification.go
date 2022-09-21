package utilities

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	modelService "github.com/tensuqiuwulu/be-service-teman-bunda/models/service"
)

func SendPushNotification(toDeviceToken string, data *modelService.NotificationData) {
	url, _ := url.Parse("https://fcm.googleapis.com/fcm/send")

	makeReqBody := modelService.PushNotificationRequestBody{
		ToDeviceToken: []string{toDeviceToken},
		Notification: modelService.NotificationData{
			Title: data.Title,
			Body:  data.Body,
		},
	}

	postBody, _ := json.Marshal(makeReqBody)

	reqBody := io.NopCloser(strings.NewReader(string(postBody)))

	req := &http.Request{
		Method: "POST",
		URL:    url,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {config.GetConfig().Fcm.Serverkey},
		},
		Body: reqBody,
	}

	reqDump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("REQUEST:\n%s", string(reqDump))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("An Error Occured %v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res.Body)
	fmt.Println(string(body))

}
