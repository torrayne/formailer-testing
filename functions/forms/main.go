package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/djatwood/formailer"
	"github.com/djatwood/formailer/handlers"
)

func main() {
	contact := formailer.Form{Name: "Contact"}
	contact.AddEmail(formailer.Email{
		To:      "daniel@atwood.io",
		From:    "daniel@atwood.io",
		Subject: "New Contact Form Submission",
	})

	formailer.Add(contact)
	lambda.Start(handlers.Netlify(formailer.DefaultConfig))
}
