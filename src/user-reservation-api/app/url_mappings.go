package app

import (
	"user-reservation-api/controller"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// mapping user
	router.POST("/user", controller.InsertUser)
	router.GET("/user/:id", controller.GetUserById)
	router.GET("/user", controller.GetUsers)
	router.POST("/login", controller.UserLogin)

	// mapping booking
	router.GET("/booking/:id", bookingController.GetBookingById)
	router.GET("/booking", bookingController.GetBookings)
	router.POST("/booking", bookingController.InsertBooking)
	router.GET("/booking/user/:user_id", bookingController.GetBookingsByUserId)
	router.GET("/hotel/availability/:id/:start_date/:end_date", bookingController.GetAvailabilityByIdAndDate)

	log.Info("Finishing mappings configurations")
}
