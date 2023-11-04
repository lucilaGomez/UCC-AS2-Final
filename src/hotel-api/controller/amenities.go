package controller

import (
	"hotel-api/dto"
	"hotel-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// InsertAmenity maneja la solicitud para insertar una nueva amenidad.
func InsertAmenity(c *gin.Context) {

	// BindJSON decodifica el JSON del cuerpo de la solicitud en la estructura AmenityDto.
	var amenityDto dto.AmenityDto
	err := c.BindJSON(&amenityDto)

	// Verificar errores en la decodificación del JSON.
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamar al servicio para insertar la amenidad y obtener el resultado.
	amenityDto, er := service.AmenityService.InsertAmenity(amenityDto)

	// Manejar errores del servicio.
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	// Responder con el resultado exitoso y la amenidad creada.
	c.JSON(http.StatusCreated, amenityDto)
}

// GetAmenities maneja la solicitud para obtener todas las amenidades.
func GetAmenities(c *gin.Context) {

	// Crear una estructura para almacenar las amenidades.
	var amenitiesDto dto.AmenitiesDto

	// Llamar al servicio para obtener todas las amenidades.
	amenitiesDto, err := service.AmenityService.GetAmenities()

	// Manejar errores del servicio.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Responder con la lista de amenidades.
	c.JSON(http.StatusOK, amenitiesDto)
}

// DeleteAmenityById maneja la solicitud para eliminar una amenidad por su ID.
func DeleteAmenityById(c *gin.Context) {

	// Obtener el parámetro "id" de la URL.
	id := c.Param("id")

	// Llamar al servicio para eliminar la amenidad por su ID.
	err := service.AmenityService.DeleteAmenityById(id)

	// Manejar errores del servicio.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Responder con un mensaje indicando que la amenidad se eliminó con éxito.
	c.JSON(http.StatusOK, gin.H{"message": "amenity deleted successfully"})
}
