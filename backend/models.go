package main

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
	BirthTime string `json:"birth_time"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
