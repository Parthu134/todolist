package service

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendOtp(email, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "partha@vananam.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	d := gomail.NewDialer("smtp.gmail.com", 587, "partha@vananam.com", "tgvw lojy dhnp qoms")
	err:=d.DialAndSend(m)
	if err!=nil{
		fmt.Println("Send Email error", err)
	}
	return err
}
