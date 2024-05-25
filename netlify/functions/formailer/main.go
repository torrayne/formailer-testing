package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/torrayne/formailer"
	"github.com/torrayne/formailer/handlers"
)

func main() {
	fmt.Println("HIT FORMAILER")
	contact := formailer.Form{Name: "Contact"}
	contact.AddEmail(formailer.Email{
		To:      "rayne@atwood.io",
		From:    "rayne@atwood.io",
		Subject: "New Contact Form Submission",
	})

	formailer.Add(&contact)
	lambda.Start(handlers.Netlify(formailer.DefaultConfig))
}
