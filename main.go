package main

import (
	"01-Login/platform/authenticator"
	"01-Login/platform/router"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Print("Server listening on http://localhost:8080/")
	if err := http.ListenAndServe("0.0.0.0:8080", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
