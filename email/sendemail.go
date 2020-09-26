package email

import (
	"net/smtp"
	"log"
)

func SendEmail(body string) {
	email 		:= getUserEmail()
	password 	:= getUserPassword()
	host 			:= getSMTPServer()
	port			:= getTLSPort()
	to				:= getUserEmail()

	auth := smtp.PlainAuth("", email, password, host)

	msg := []byte(	"From: " + email + "\n" +
									"To: " + to + "\n" +
									"Subject: Birthdays Today" + "\n\n" +
									body)

	err := smtp.SendMail(host+":"+port, auth, email, []string{to}, msg)
	if err != nil {
		log.Fatal(err)
	}
}
