package main

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"os/exec"
	"strings"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
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

	var err error
	db, err = sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	// Set the maximum number of open connections to the database.
	db.SetMaxOpenConns(10)
}

// Function to fetch user data from the database
func fetchUserData(id int) (User, error) {
	// SQL query to fetch the user data
	query := `SELECT id, first_name, birth_date, birth_time, city, natal_chart FROM profile WHERE id = $1`

	// Execute the query
	row := db.QueryRow(query, id)

	// Scan the result into a User struct
	var user User
	err := row.Scan(&user.ID, &user.FirstName, &user.BirthDate, &user.BirthTime, &user.City, &user.NatalChartPath)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// Function to update user data in the database
func updateUserData(id int, natalChartPath string) error {
	// SQL query to update the user data
	query := `UPDATE profile SET natal_chart = $1 WHERE id = $2`

	// Execute the query
	_, err := db.Exec(query, natalChartPath, id)
	if err != nil {
		return err
	}

	return nil
}

// Function to call the Python script to generate the natal chart
func callPythonScript(user User, chartType string) (string, error) {
	// Prepare the command to call the Python script
	cmd := exec.Command("python3", "chart-generator.py", user.FirstName, user.BirthDate, user.BirthTime, user.City, chartType)

	// Capture the stdout output
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Split the output into lines
	lines := strings.Split(out.String(), "\n")

	// Initialize the SVG file path
	var svgFilePath string

	// Iterate over the lines to find the SVG file path
	for _, line := range lines {
		// If the line contains the SVG file path, save it
		if strings.Contains(line, "assets/charts/") {
			svgFilePath = strings.TrimSpace(line)
		}
	}

	// If the SVG file path was not found, return an error
	if svgFilePath == "" {
		return "", errors.New("SVG file path not found in Python script output")
	}

	// Return the SVG file path and no error
	return svgFilePath, nil
}
