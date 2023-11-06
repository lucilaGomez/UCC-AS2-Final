package service

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"hotel-api/client"
	"hotel-api/dto"
	"hotel-api/model"
	"hotel-api/queue"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id string) (dto.HotelDto, error)
	GetHotels() (dto.HotelsDto, error)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error)
	DeleteHotel(id string) error
	UpdateHotel(hotelDto dto.HotelDto) (dto.HotelDto, error)
}

var HotelService hotelServiceInterface

// Función de inicialización que asigna la implementación del servicio
func init() {
	HotelService = &hotelService{}
}

// Función para insertar un hotel
func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {
	var hotel model.Hotel
	var hotelNew model.Hotel

	// TODO buscar HOTEL EN AMADEUS
	var amadeusDTO = AmadeusService.getHotelByCity(hotelDto.City)
	log.Info(amadeusDTO)

	var idAmadeus = "el id de amadeus"

	hotel.Name = hotelDto.Name
	hotel.Description = hotelDto.Description
	hotel.RoomAmount = hotelDto.RoomAmount
	hotel.City = hotelDto.City
	hotel.StreetName = hotelDto.StreetName
	hotel.StreetNumber = hotelDto.StreetNumber
	hotel.Rate = hotelDto.Rate

	for _, amenityName := range hotelDto.Amenities {
		amenity := client.AmenityClient.GetAmenityByName(amenityName)

		if amenity.Id.Hex() == "000000000000000000000000" {
			return hotelDto, errors.New("amenity not found")
		}

		hotel.Amenities = append(hotel.Amenities, amenity)

	}

	hotelNew = client.HotelClient.InsertHotel(hotel)

	hotelDto.Id = hotelNew.Id.Hex()

	if hotelNew.Id.Hex() == "000000000000000000000000" {
		return hotelDto, errors.New("error creating hotel")
	}

	// TODO ACTUALIZAR HOTEL CON EL ID DE AMADEUS

	hotelNew.IdAmadeus = idAmadeus
	hotelDto.IdAmadeus = idAmadeus
	s.UpdateHotel(hotelDto)

	body := map[string]interface{}{
		"id":            hotelNew.Id.Hex(),
		"message":       "create",
		"name":          hotelNew.Name,
		"description":   hotelNew.Description,
		"room_amount":   hotelNew.RoomAmount,
		"city":          hotelNew.City,
		"street_name":   hotelNew.StreetName,
		"street_number": hotelNew.StreetNumber,
		"rate":          hotelNew.Rate,
		"id_amadeus":    hotelNew.IdAmadeus,
	}

	jsonBody, _ := json.Marshal(body)

	err := queue.Publish(jsonBody)

	if err != nil {
		return hotelDto, err
	}

	return hotelDto, nil
}

// Función para obtener todos los hoteles
func (s *hotelService) GetHotels() (dto.HotelsDto, error) {

	var hotels model.Hotels = client.HotelClient.GetHotels()
	var hotelsDto dto.HotelsDto

	for _, hotel := range hotels {
		var hotelDto dto.HotelDto
		hotelDto.Id = hotel.Id.Hex()
		hotelDto.Name = hotel.Name
		hotelDto.RoomAmount = hotel.RoomAmount
		hotelDto.City = hotel.City
		hotelDto.Description = hotel.Description
		hotelDto.StreetName = hotel.StreetName
		hotelDto.StreetNumber = hotel.StreetNumber
		hotelDto.Rate = hotel.Rate

		// Append thumbnail (first image)
		if len(hotel.Images) > 0 {
			hotelDto.Images = append(hotelDto.Images, hotel.Images[0])
		}

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

// Función para obtener un hotel por su ID
func (s *hotelService) GetHotelById(id string) (dto.HotelDto, error) {

	var hotel model.Hotel = client.HotelClient.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.Id.Hex() == "000000000000000000000000" {
		return hotelDto, errors.New("hotel not found")
	}
	hotelDto.Id = hotel.Id.Hex()
	hotelDto.Name = hotel.Name
	hotelDto.RoomAmount = hotel.RoomAmount
	hotelDto.Description = hotel.Description
	hotelDto.City = hotel.City
	hotelDto.StreetName = hotel.StreetName
	hotelDto.StreetNumber = hotel.StreetNumber
	hotelDto.Rate = hotel.Rate

	for _, amenity := range hotel.Amenities {
		hotelDto.Amenities = append(hotelDto.Amenities, amenity.Name)
	}

	for _, image := range hotel.Images {
		hotelDto.Images = append(hotelDto.Images, image)
	}

	return hotelDto, nil
}

// Función para eliminar un hotel por su ID
func (s *hotelService) DeleteHotel(id string) error {

	hotel := client.HotelClient.GetHotelById(id)

	if hotel.Id.Hex() == "000000000000000000000000" {
		return errors.New("hotel not found")
	}

	err := client.HotelClient.DeleteHotelById(id)

	return err
}

// Función para actualizar un hotel
func (s *hotelService) UpdateHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {

	hotel := client.HotelClient.GetHotelById(hotelDto.Id)

	if hotel.Id.Hex() == "000000000000000000000000" {
		return hotelDto, errors.New("hotel not found")
	}

	hotel.Name = hotelDto.Name
	hotel.City = hotelDto.City
	hotel.StreetName = hotelDto.StreetName
	hotel.StreetNumber = hotelDto.StreetNumber
	hotel.Rate = hotelDto.Rate
	hotel.Description = hotelDto.Description
	hotel.RoomAmount = hotelDto.RoomAmount
	hotel.Amenities = model.Amenities{}

	for _, amenityName := range hotelDto.Amenities {
		amenity := client.AmenityClient.GetAmenityByName(amenityName)

		if amenity.Id.Hex() == "000000000000000000000000" {
			return hotelDto, errors.New("amenity not found")
		}

		hotel.Amenities = append(hotel.Amenities, amenity)
	}

	for _, image := range hotelDto.Images {
		hotel.Images = append(hotel.Images, image)
	}

	hotel = client.HotelClient.UpdateHotelById(hotel)

	if hotel.Id.Hex() == "000000000000000000000000" {
		return hotelDto, errors.New("error updating hotel")
	}

	body := map[string]interface{}{
		"id":            hotel.Id.Hex(),
		"message":       "update",
		"name":          hotel.Name,
		"description":   hotel.Description,
		"room_amount":   hotel.RoomAmount,
		"city":          hotel.City,
		"street_name":   hotel.StreetName,
		"street_number": hotel.StreetNumber,
		"rate":          hotel.Rate,
		"id_amadeus":    hotel.IdAmadeus,
	}

	jsonBody, _ := json.Marshal(body)

	err := queue.Publish(jsonBody)

	if err != nil {
		return hotelDto, err
	}

	return hotelDto, nil

}
