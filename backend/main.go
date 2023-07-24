package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize the database connection
	initDB()
	log.Println("Database connected")

	// Create a new router
	router := mux.NewRouter()

	// Register the API endpoints
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")

	// Start the HTTP server
	cors := cors.New(cors.Options{
		AllowedHeaders: []string{"Content-Type", "Authorization", "X-Requested-With", "Content-Length", "Accept", "Origin"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", cors.Handler(router))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
