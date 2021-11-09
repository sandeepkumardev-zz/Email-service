package utils

import (
	"log"
	"net/smtp"
	"os"
)

func SendEmail(to []string, message []byte, resChan chan string, errChan chan string) {
	// Create authentication
	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), os.Getenv("SMTP_HOST"))
	log.Println("sending email...")
	// Send message
	err := smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, os.Getenv("EMAIL"), to, message)
	if err != nil {
		errChan <- "Message is not sent to smtp server!"
	}
	log.Println("Message sent to smtp server")
	resChan <- "Message sent."
}
