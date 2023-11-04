package app

import (
	"hotel-api/controller"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Mapeo de rutas relacionadas con hoteles
	router.POST("/hotel", controller.InsertHotel)
	router.GET("/hotel/:id", controller.GetHotelById)
	router.GET("/hotel", controller.GetHotels)
	router.POST("/hotel/:id/images", controller.InsertImages)
	router.DELETE("/hotel/:id", controller.DeleteHotel)
	router.PUT("/hotel/:id", controller.UpdateHotel)

	// Mapeo de rutas relacionadas con amenidades
	router.POST("/amenity", controller.InsertAmenity)
	router.GET("/amenity", controller.GetAmenities)
	router.DELETE("amenity/:id", controller.DeleteAmenityById)

	log.Info("Finishing mappings configurations")
}
