package utilities

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
)

type SendWaRequest struct {
	ToNumber        string     `json:"to_number"`
	ToName          string     `json:"to_name"`
	MssgTemplateId  string     `json:"message_template_id"`
	ChIntegrationId string     `json:"channel_integration_id"`
	Language        Language   `json:"language"`
	Parameters      Parameters `json:"parameters"`
}

type Language struct {
	Code string `json:"code"`
}

type Parameters struct {
	Bodys []Body `json:"body"`
}

type Body struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	ValueText string `json:"value_text"`
}

func SendWhatsapp(toNumber string, toName string, data *Body, mssgTemplateId string) {
	url, _ := url.Parse(string(config.GetConfig().Whatsapp.WhatsappUrl))

	makeReqBody := SendWaRequest{
		ToNumber:        toNumber,
		ToName:          toName,
		MssgTemplateId:  mssgTemplateId,
		ChIntegrationId: config.GetConfig().Whatsapp.ChannelId,
		Language: Language{
			Code: "id",
		},
		Parameters: Parameters{
			Bodys: []Body{
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
		log.Fatalf("An Error Occured %v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res.Body)
	fmt.Println(string(body))
}
