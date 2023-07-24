package main

import (
	"database/sql"
	"fmt"
	"net/url"

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
