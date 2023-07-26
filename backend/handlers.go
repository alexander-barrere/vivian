package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// login handles the login API endpoint.
func login(w http.ResponseWriter, r *http.Request) {
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

	if user.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Password is required")
		return
	}

	// Check if email address and password match a user in the database
	var hashedPassword, firstName, lastName, city, state, country, birthDate, birthTime, latitude, longitude string
	query := "SELECT id, first_name, last_name, password, city, state, country, birth_date, birth_time, latitude, longitude FROM profile WHERE email = $1"
	err = db.QueryRow(query, user.Email).Scan(&user.ID, &firstName, &lastName, &hashedPassword, &city, &state, &country, &birthDate, &birthTime, &latitude, &longitude)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error checking for existing profile: %v", err))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid email address or password")
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.ID,
		"email":      user.Email,
		"first_name": firstName,
		"last_name":  lastName,
		"city":       city,
		"state":      state,
		"country":    country,
		"birth_date": birthDate,
		"birth_time": birthTime,
		"latitude":   latitude,
		"longitude":  longitude,
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("5f82737a11461f864f9417395fde341921fedf1b980fc907ed1f6e6efca01064"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error signing token: %v", err))
		return
	}

	// Return the token in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

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

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error hashing password: %v", err))
		return
	}

	// Insert user data and latitude/longitude values into the `profile` table
	query = "INSERT INTO profile (first_name, last_name, email, password, birth_date, birth_time, city, state, country, latitude, longitude) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	err = db.QueryRow(query, user.FirstName, user.LastName, user.Email, string(hashedPassword), user.BirthDate, user.BirthTime, user.City, user.State, user.Country, latitudeStr, longitudeStr).Scan(&user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error inserting data: %v", err))
		return
	}

	// Generate the Natal chart for the user and store the SVG file path
	fmt.Println("Generating Natal chart...")
	svgPath, err := callPythonScript(user, "Natal")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error generating Natal chart: %v", err))
		return
	}
	user.NatalChartPath = svgPath

	// Update the user data in the database
	fmt.Println("Updating user data in the database...")
	err = updateUserData(user.ID, user.NatalChartPath)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error updating user data: %v", err))
		return
	}

	// Set the latitude and longitude values for the response
	user.Latitude = latitudeStr
	user.Longitude = longitudeStr

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	return
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameters
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Fetch the user data from the database
	user, err := fetchUserData(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching user data: %v", err))
		return
	}

	// Return the user data in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func validateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			token, err := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("5f82737a11461f864f9417395fde341921fedf1b980fc907ed1f6e6efca01064"), nil
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

func protectedEndpointHandler(w http.ResponseWriter, r *http.Request) {
	// This is a protected endpoint, handle it here
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Access to protected endpoint successful")
}

// respondWithError writes an error message to the response.
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
}
