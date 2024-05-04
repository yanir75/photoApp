package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"update_service/platform/authenticator"
	"update_service/platform/router"
)

func setupRoutes() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Print("Server listening on http://localhost:8080/")

	http.ListenAndServe(":8080", rtr)
}

func main() {
	fmt.Println("Setting up server")
	setupRoutes()
}
