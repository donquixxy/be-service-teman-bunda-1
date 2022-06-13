package utilities

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"gopkg.in/gomail.v2"
)

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return "", err
	}
	return buf.String(), nil
}

func SendEmail(to string, subject string, data interface{}, templateFile string) error {
	result, _ := ParseTemplate(templateFile, data)
	m := gomail.NewMessage()
	m.SetHeader("From", string(config.GetConfig().Email.FromEmail))
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", result)
	// m.Attach(templateFile) // attach whatever you want
	senderPort := 465
	d := gomail.NewDialer("smtp.gmail.com", senderPort, string(config.GetConfig().Email.FromEmail), string(config.GetConfig().Email.FromEmailPassword))
	err := d.DialAndSend(m)
	if err != nil {
		// panic(err)
		fmt.Println("Error = ", err)
	}
	return err
}
