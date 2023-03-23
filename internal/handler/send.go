package handler

import (
	"email-api/private"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data from the POST request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Get the fields from the form data
	from := r.Form.Get("from")
	to := r.Form.Get("to")
	subject := r.Form.Get("subject")
	body := r.Form.Get("body")

	// Send the email
	err = sendMail(from, to, subject, body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return a success message
	fmt.Fprintln(w, "Email sent successfully!")
}

func sendMail(from, to, subject, body string) error {
	// Set up the email message
	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// Set up the SMTP server information
	config := private.GetConfig() //change this to ENV variables at some point
	smtpServer := config.Server
	auth := smtp.PlainAuth("", config.User, config.Password, smtpServer)

	// Send the email using SMTP
	err := smtp.SendMail(smtpServer+":587", auth, from, []string{to}, message)
	if err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	return nil
}
