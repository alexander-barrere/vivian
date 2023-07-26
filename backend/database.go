package main

import (
	"database/sql"
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

func callPythonScript(user User, chartType string) (string, error) {
	fmt.Printf("Running Python script to generate %s chart...\n", chartType)
	cmd := exec.Command("python3", "chart-generator.py", user.FirstName, user.BirthDate, user.BirthTime, user.City, chartType)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	fmt.Printf("Python script output: %s\n", output)
	return strings.TrimSpace(string(output)), nil
}
