package api

import (
	"log"
	"net/http"

	"github.com/djatwood/formailer"
	"github.com/djatwood/formailer/handlers"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	contact := formailer.Form{Name: "Contact"}
	contact.AddEmail(formailer.Email{
		To:      "daniel@atwood.io",
		From:    "daniel@atwood.io",
		Subject: "New Contact Form Submission",
	})

	formailer.Add(contact)
	log.Print(formailer.DefaultConfig)
	handlers.Vercel(formailer.DefaultConfig, w, r)
}
