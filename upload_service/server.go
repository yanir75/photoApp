package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/joho/godotenv"


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

	log.Print("Server listening on http://localhost:3000/")
	
	http.ListenAndServe(":3000", rtr)
}

func main() {
	fmt.Println("Setting up server")
	setupRoutes()
}
