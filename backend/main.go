package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/rubenv/opencagedata"
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	BirthDate string `json:"birth_date"`
	BirthTime string `json:"birth_time"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
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
	log.Println("Database connected") // Added log statement

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

	log.Println("Starting server on port 8080") // Added log statement
	err := http.ListenAndServe(":8080", cors.Handler(router))
	if err != nil {
		log.Fatalf("Error starting server: %v", err) // Added log statement
	}
}

func initDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
}

func geocode(city, state, country string) (string, string, error) {
	geocoder := opencagedata.NewGeocoder("89d5a7e1287b40b9b8418e3e7775e054")

	query := fmt.Sprintf("%s, %s, %s", city, state, country)
	result, err := geocoder.Geocode(query, nil)
	if err != nil {
		return "", "", err
	}

	if len(result.Results) > 0 {
		f_result := result.Results[0]
		latitude := fmt.Sprintf("%.7f", f_result.Geometry.Latitude)
		longitude := fmt.Sprintf("%.7f", f_result.Geometry.Longitude)
		return latitude, longitude, nil
	}

	return "", "", fmt.Errorf("No results found for query: %s", query)
}

func register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding request: %v", err)
		return
	}

	// Validate user input
	if user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Email address is required")
		return
	}

	if !isValidEmail(user.Email) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid email address")
		return
	}

	if user.City == "" || user.State == "" || user.Country == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "City, state, and country are required")
		return
	}

	// Perform geocode lookup
	latitudeStr, longitudeStr, err := geocode(user.City, user.State, user.Country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Geocoding failed",
			"error":   fmt.Sprintf("Error geocoding: %v", err),
		})
		return
	}

	// Check if email address is already in use
	query := "SELECT COUNT(*) FROM profile WHERE email = $1"
	var count int
	err = db.QueryRow(query, user.Email).Scan(&count)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error checking for existing profile: %v", err)
		return
	}

	if count > 0 {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "A profile with that email address already exists")
		return
	}

	// Insert user data and latitude/longitude values into the `profile` table
	query = "INSERT INTO profile (first_name, last_name, email, password, birth_date, birth_time, city, state, country, latitude, longitude) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	err = db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.BirthDate, user.BirthTime, user.City, user.State, user.Country, latitudeStr, longitudeStr).Scan(&user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error inserting data: %v", err)
		return
	}

	// Set the latitude and longitude values for the response
	user.Latitude = latitudeStr
	user.Longitude = longitudeStr

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	return
}

func isValidEmail(email string) bool {
	if email == "" {
		return false
	}

	// Use regex to validate email format
	// ...

	return true
}
