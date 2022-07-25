package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
)

func SendSmsOtp(to string, otpCode string) {

	fmt.Println("REQUEST KIRIM PESAN !!!!")
	message := fmt.Sprintf("Kode Verifikasi Teman Bunda Anda adalah: %s *JANGAN BERIKAN KODE INI KEPADA SIAPAPUN, TERMASUK PIHAK TEMAN BUNDA* Hubungi 081228512244 untuk bantuan.", otpCode)

	postBody, _ := json.Marshal(map[string]string{
		"userkey": config.GetConfig().Sms.UserKey,
		"passkey": config.GetConfig().Sms.PassKey,
		"to":      to,
		"message": message,
	})

	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("https://console.zenziva.net/wareguler/api/sendWA/", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	fmt.Printf(sb)
}
