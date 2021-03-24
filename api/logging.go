package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Logger(w http.ResponseWriter, r *http.Request) {
	warn := log.New(os.Stdout, "warn: ", log.Ldate|log.Ltime)
	warn.Println(errors.New("This is a warning"))
	er := log.New(os.Stderr, "error: ", log.Ldate|log.Ltime)
	er.Println(errors.New("This is an error"))
	fmt.Fprintln(w, "Hello, World!")
}
