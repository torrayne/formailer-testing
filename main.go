package main

import (
	"net/http"

	"github.com/djatwood/formailer"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	fs := http.Dir("./public")
	http.HandleFunc("/api/", handle)
	http.Handle("/", http.FileServer(fs))
	http.ListenAndServe(":4120", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	cfg := make(formailer.Config)
	cfg.Set(&formailer.Form{
		Name:    "Contact",
		To:      "daniel@atwood.io",
		From:    "daniel@atwood.io",
		Subject: "New Contact Form Submission",
	})

	cfg.Vercel(w, r)
}
