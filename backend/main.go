package main

import (
	"log"
	"net/http"
	"strconv"

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
	router.Handle("/protected-endpoint", validateTokenMiddleware(protectedEndpointHandler)).Methods("GET")
	router.HandleFunc("/natal-chart/{id}", generateNatalChart).Methods("GET")

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

// New function to handle the /natal-chart/{id} route
func generateNatalChart(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the route parameters
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Fetch the user data from the database
	user, err := fetchUserData(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Pass the user data to the Python script and get the file path of the Natal chart SVG
	svgPath, err := callPythonScript(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store the SVG file path in the database
	err = updateUserData(id, svgPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Natal chart generated successfully"))
}
