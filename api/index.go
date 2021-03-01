package api

import (
	"net/http"

	"github.com/djatwood/formailer"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	cfg := make(formailer.Config)
	cfg.Set(&formailer.Form{
		Name:    "Contact",
		To:      "datwood@wildwoodhealth.org",
		From:    "donotreply@wildwoodhealth.org",
		Subject: "New Contact Form Submission",
	})

	cfg.Vercel(w, r)
}
