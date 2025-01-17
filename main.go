package main

import (
	"net/http"
	"os"

	"meli-api/controller"
	"meli-api/repository"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// Load the .env file
	var err error = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Check if the required environment variables are set
	if os.Getenv("SHORT_URL_DOMAIN") == "" {
		panic("SHORT_URL_DOMAIN is required")
	}

	// Connect to the database
	err = repository.Connect()
	if err != nil {
		panic("Error connecting to the database " + err.Error())
	}

	// Create the router
	r := mux.NewRouter()
	r.HandleFunc("/", controller.PostHandler).Methods("POST")
	r.HandleFunc("/{key}", controller.PatchHandler).Methods("PATCH")
	r.HandleFunc("/{key}", controller.DeleteHandler).Methods("DELETE")
	r.HandleFunc("/key/{key}", controller.GetHandler).Methods("GET")
	r.HandleFunc("/all", controller.GetAllHandler).Methods("GET")
	r.HandleFunc("/healthcheck", controller.HealthCheckHandler).Methods("GET")

	// Get the HTTP host and port
	httpHost := os.Getenv("HTTP_HOST")
	httpPort := os.Getenv("HTTP_PORT")

	// Start the HTTP server
	http.ListenAndServe(httpHost+":"+httpPort, r)
}
