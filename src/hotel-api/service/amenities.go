package service

import (
	"errors"
	"hotel-api/client"
	"hotel-api/dto"
	"hotel-api/model"
)

type amenityService struct{}

type amenityServiceInterface interface {
	InsertAmenity(amenityDto dto.AmenityDto) (dto.AmenityDto, error)
	GetAmenities() (dto.AmenitiesDto, error)
	DeleteAmenityById(id string) error
}

var AmenityService amenityServiceInterface

func init() {
	AmenityService = &amenityService{}
}

func (s *amenityService) InsertAmenity(amenityDto dto.AmenityDto) (dto.AmenityDto, error) {

	// Insertar una nueva amenidad.
	checkAmenity := client.AmenityClient.GetAmenityByName(amenityDto.Name)
	// Verificar si ya existe una amenidad con el mismo nombre. Si ya existe,
	// se devuelve un error indicando que la amenidad ya existe.

	if checkAmenity.Id.Hex() == "000000000000000000000000" {
		var amenity model.Amenity

		amenity.Name = amenityDto.Name

		amenity = client.AmenityClient.InsertAmenity(amenity)

		if amenity.Id.Hex() == "000000000000000000000000" {
			return amenityDto, errors.New("error creating amenity")
		}

		amenityDto.Id = amenity.Id.Hex()

		return amenityDto, nil

	}

	return amenityDto, errors.New("amenity already exists")

}

func (s *amenityService) GetAmenities() (dto.AmenitiesDto, error) {
	// Se encarga de obtener todas las amenidades disponibles.

	var amenities model.Amenities = client.AmenityClient.GetAmenities()
	// Para obtener la lista de todas las amenidades desde el cliente de amenidades.
	var amenitiesDto dto.AmenitiesDto

	for _, amenity := range amenities {
		var amenityDto dto.AmenityDto
		amenityDto.Id = amenity.Id.Hex()
		amenityDto.Name = amenity.Name

		amenitiesDto = append(amenitiesDto, amenityDto)
	}

	return amenitiesDto, nil
}

func (s *amenityService) DeleteAmenityById(id string) error {
	// Se encarga de eliminar una amenidad según su ID.

	err := client.AmenityClient.DeleteAmenityById(id)
	//  Para eliminar la amenidad mediante el cliente de amenidades.

	if err != nil {
		return errors.New("failed to delete amenity")
	}

	return nil
}
