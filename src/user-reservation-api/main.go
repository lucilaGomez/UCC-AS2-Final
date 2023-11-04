package main

import (
	"user-reservation-api/app"
	"user-reservation-api/db"
)

func main() {

	db.StartDbEngine()
	app.StartRoute()
}
