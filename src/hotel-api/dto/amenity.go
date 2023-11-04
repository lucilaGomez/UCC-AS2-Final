package dto

type AmenityDto struct {
	Id   string `json:"id"`
	Name string `json:"name" validate:"required"`
}

type AmenitiesDto []AmenityDto
