package controller

import (
	"fmt"
	"hotel-api/dto"
	"hotel-api/service"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// InsertHotel maneja la solicitud para insertar un nuevo hotel.
func InsertHotel(c *gin.Context) {

	// BindJSON decodifica el JSON del cuerpo de la solicitud en la estructura HotelDto.
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	// Verificar errores en la decodificación del JSON.
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamar al servicio para insertar el hotel y obtener el resultado.v
	hotelDto, er := service.HotelService.InsertHotel(hotelDto)

	// Manejar errores del servicio.
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	// Responder con el resultado exitoso y el hotel creado.
	c.JSON(http.StatusCreated, hotelDto)
}

// GetHotelById maneja la solicitud para obtener un hotel por su ID.
func GetHotelById(c *gin.Context) {

	// Obtener el parámetro "id" de la URL.
	id := c.Param("id")
	var hotelDto dto.HotelDto

	// Llamar al servicio para obtener el hotel por su ID.
	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Responder con la información del hotel.
	c.JSON(http.StatusOK, hotelDto)
}

// GetHotels maneja la solicitud para obtener todos los hoteles.
func GetHotels(c *gin.Context) {

	// estructura para almacenar la lista de hoteles.
	var hotelsDto dto.HotelsDto

	// Llamar al servicio
	hotelsDto, err := service.HotelService.GetHotels()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
}

// DeleteHotel maneja la solicitud para eliminar un hotel por su ID.
func DeleteHotel(c *gin.Context) {

	// Obtener el parámetro "id" de la URL.
	id := c.Param("id")

	err := service.HotelService.DeleteHotel(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted"})
}

// UpdateHotel maneja la solicitud para actualizar un hotel.
func UpdateHotel(c *gin.Context) {

	// Obtener el parámetro "id" de la URL.
	id := c.Param("id")
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	// Verificar errores en la decodificación del JSON.
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Asignar el ID al DTO antes de llamar al servicio.
	hotelDto.Id = id

	// Llamar al servicio para actualizar el hotel.
	hotelDto, err = service.HotelService.UpdateHotel(hotelDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotelDto)
}

// InsertImages maneja la solicitud para insertar imágenes en un hotel.
func InsertImages(c *gin.Context) {

	// Crea un DTO para el hotel
	var hotelDto dto.HotelDto

	// Obtener el ID del hotel de los parámetros de la URL.
	id := c.Param("id")

	// Llamar al servicio para obtener la información del hotel por su ID.
	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtener los archivos enviados en la solicitud.
	form, _ := c.MultipartForm()
	files := form.File["images"]

	// Obtener el número actual de imágenes en el hotel.
	imageCount := len(hotelDto.Images)

	// Iterar sobre los archivos y guardarlos en el sistema de archivos.
	for i, file := range files {

		fileExt := path.Ext(file.Filename)

		//Filename as [hotel_id]-[image_number].[file_extension]
		fileName := fmt.Sprintf("%s-%d%s", id, i+1+imageCount, fileExt)

		// Guardar el archivo en el sistema de archivos.
		err = c.SaveUploadedFile(file, "Images/"+fileName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		hotelDto.Images = append(hotelDto.Images, fileName)
	}

	// Llamar al servicio para actualizar la información del hotel con las nuevas imágenes
	hotelDto, err = service.HotelService.UpdateHotel(hotelDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotelDto)
}
