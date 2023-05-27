package helper

import (
	"gopkg.in/gomail.v2"
	"log"
)

func SendEmail(subject string, email string, message string) {
	const CONFIG_SMTP_HOST = "smtp.titan.email"
	const CONFIG_SMTP_PORT = 465
	const CONFIG_SENDER_NAME = "MyFin <help@hellomyfin.engineer>"
	const CONFIG_AUTH_EMAIL = "help@hellomyfin.engineer"
	const CONFIG_AUTH_PASSWORD = "adminmyfin"

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)
	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

}
