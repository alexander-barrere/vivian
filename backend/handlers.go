package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// login handles the login API endpoint.
func login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding request: %v", err))
		return
	}

	if count == 0 {
		respondWithError(w, http.StatusUnauthorized, "Invalid email address or password")
		return
	}

	// Validate user input
	if user.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Email address is required")
		return
	}

	if !isValidEmail(user.Email) {
		respondWithError(w, http.StatusBadRequest, "Invalid email address")
		return
	}

	if user.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Password is required")
		return
	}

	// Check if email address and password match a user in the database
	query := "SELECT COUNT(*) FROM profile WHERE email = $1 AND password = $2"
	var count int
	err = db.QueryRow(query, user.Email, user.Password).Scan(&count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error checking for existing profile: %v", err))
		return
	}

	if count == 0 {
		respondWithError(w, http.StatusUnauthorized, "Invalid email address or password")
		return
	}

	// If the email and password match, return a success status
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Login successful")
	return
}

// register handles the register API endpoint.
func register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding request: %v", err))
		return
	}

	// Validate user input
	if user.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Email address is required")
		return
	}

	if !isValidEmail(user.Email) {
		respondWithError(w, http.StatusBadRequest, "Invalid email address")
		return
	}

	if user.City == "" || user.State == "" || user.Country == "" {
		respondWithError(w, http.StatusBadRequest, "City, state, and country are required")
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
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error checking for existing profile: %v", err))
		return
	}

	if count > 0 {
		respondWithError(w, http.StatusConflict, "A profile with that email address already exists")
		return
	}

	// Insert user data and latitude/longitude values into the `profile` table
	query = "INSERT INTO profile (first_name, last_name, email, password, birth_date, birth_time, city, state, country, latitude, longitude) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	err = db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.BirthDate, user.BirthTime, user.City, user.State, user.Country, latitudeStr, longitudeStr).Scan(&user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error inserting data: %v", err))
		return
	}

	// Set the latitude and longitude values for the response
	user.Latitude = latitudeStr
	user.Longitude = longitudeStr

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	return
}

func validateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			token, err := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("your-secret-key"), nil
			})
			if err != nil {
				respondWithError(w, http.StatusUnauthorized, "Invalid token")
				return
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// Add claims to request context
				ctx := context.WithValue(r.Context(), "email", claims["email"])
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				respondWithError(w, http.StatusUnauthorized, "Invalid token")
				return
			}
		} else {
			respondWithError(w, http.StatusUnauthorized, "An authorization header is required")
			return
		}
	})
}

// respondWithError writes an error message to the response.
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
}
