package main

import (
	"email-api/internal/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/send-mail", handler.SendMailHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
