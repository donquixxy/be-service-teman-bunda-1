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

func SendWhatsapp(toNumber string, toName string, data *modelService.WhatsappBody, mssgTemplateId string) {
	url, _ := url.Parse(string(config.GetConfig().Whatsapp.WhatsappUrl))

	makeReqBody := modelService.SendWhatsappRequest{
		ToNumber:        toNumber,
		ToName:          toName,
		MssgTemplateId:  mssgTemplateId,
		ChIntegrationId: config.GetConfig().Whatsapp.ChannelId,
		Language: modelService.WhatsappLanguage{
			Code: "id",
		},
		Parameters: modelService.WhatsappParameters{
			Bodys: []modelService.WhatsappBody{
				{
					Key:       data.Key,
					Value:     data.Value,
					ValueText: data.ValueText,
				},
			},
		},
	}
	postBody, _ := json.Marshal(makeReqBody)

	reqBody := io.NopCloser(strings.NewReader(string(postBody)))

	req := &http.Request{
		Method: "POST",
		URL:    url,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {"bearer " + config.GetConfig().Whatsapp.WhatsappToken},
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
