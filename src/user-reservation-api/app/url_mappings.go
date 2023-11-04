package app

import (
	"user-reservation-api/controller"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Add all methods and its mappings
	router.POST("/user", controller.InsertUser)
	router.GET("/user/:id", controller.GetUserById)
	router.GET("/user", controller.GetUsers)
	router.POST("/login", controller.UserLogin)

	log.Info("Finishing mappings configurations")
}
