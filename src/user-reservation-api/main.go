package main

import (
	"fmt"
	"user-reservation-api/app"
	"user-reservation-api/cache"
	"user-reservation-api/db"
)

func main() {

	db.StartDbEngine()

	fmt.Println("Iniciando cache")
	cache.Init_cache()

	app.StartRoute()
}
