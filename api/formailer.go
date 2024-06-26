package api

import (
	"net/http"

	"github.com/torrayne/formailer"
	"github.com/torrayne/formailer/handlers"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	contact := formailer.New("Contact")
	contact.AddEmail(formailer.Email{
		To:      "rayne@atwood.io",
		From:    "rayne@atwood.io",
		Subject: "New Contact Form Submission",
	})

	handlers.Vercel(formailer.DefaultConfig, w, r)
}
