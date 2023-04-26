package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type User struct {
	ID        int64   `json:"id"`
	Email     string  `json:"email"`
	Password  string  `json:"-"`
	BirthDate string  `json:"birth_date"`
	BirthTime string  `json:"birth_time"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var db *sql.DB

func main() {
	// Define the database connection details
	dbUser := "starfja8_vivian"
	dbPass := "PZq(tDO^0NjV" // Your actual password
	dbName := "starfja8_users"
	dbHost := "localhost"
	dbPort := "5432"
	dbSSLMode := "disable"

	// URL-encode the password
	encodedPass := url.QueryEscape(dbPass)

	// Construct the database connection string
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, encodedPass, dbHost, dbPort, dbName, dbSSLMode)

	// Initialize the database connection
	initDB(dbURI)

	// Create a new router
	router := mux.NewRouter()

	// Register the API endpoints
	router.HandleFunc("/register", register).Methods("POST")

	// Start the HTTP server
	cors := cors.New(cors.Options{
		AllowedHeaders: []string{"Content-Type", "Authorization", "X-Requested-With", "Content-Length", "Accept", "Origin"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	http.ListenAndServe(":8080", cors.Handler(router))
} // <- Added closing brace here

func initDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	// Perform input validation, hashing password, and other business logic here

	query := "INSERT INTO profile (email, password, birth_date, birth_time, city, state, country, latitude, longitude) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	err := db.QueryRow(query, user.Email, user.Password, user.BirthDate, user.BirthTime, user.City, user.State, user.Country, user.Latitude, user.Longitude).Scan(&user.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error registering user: %v", err) // Update this line to include the error details
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
