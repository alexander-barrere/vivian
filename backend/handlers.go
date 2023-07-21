package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
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

	if user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Password is required")
		return
	}

	// Check if email address and password match a user in the database
	query := "SELECT COUNT(*) FROM profile WHERE email = $1 AND password = $2"
	var count int
	err = db.QueryRow(query, user.Email, user.Password).Scan(&count)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error checking for existing profile: %v", err)
		return
	}

	if count == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid email address or password")
		return
	}

	// If the email and password match, return a success status
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Login successful")
	return
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
