package utils

import (
	"email/models"
	"encoding/json"
	"log"
	"net/smtp"
	"os"

	"github.com/adjust/rmq/v4"
	logrus "github.com/sirupsen/logrus"
	// "github.com/matcornic/hermes/v2"
)

func SendEmail(delivery rmq.Delivery) {
	// var task interface{}
	var T models.EmailTemplate
	json.Unmarshal([]byte(delivery.Payload()), &T)

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + T.Subject + "!\n"
	msg := []byte(subject + mime + "\n" + T.Body)

	// Create authentication
	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), os.Getenv("SMTP_HOST"))
	log.Println("sending email...")
	// Send message
	err := smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, os.Getenv("EMAIL"), T.To, []byte(msg))
	if err != nil {
		logrus.Error("Message is not sent to smtp server!")
	}
	log.Println("Message sent to smtp server")
	logrus.Info("Message sent.")

	delivery.Ack()
}
