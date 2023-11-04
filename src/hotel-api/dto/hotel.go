package dto

type HotelDto struct {
	Id           string   `json:"id"`
	Name         string   `json:"name" validate:"required"`
	RoomAmount   int      `json:"room_amount" validate:"required"`
	Description  string   `json:"description" validate:"required"`
	City         string   `json:"city" validate:"required"`
	StreetName   string   `json:"street_name" validate:"required"`
	StreetNumber int      `json:"street_number" validate:"required"`
	Rate         float64  `json:"rate" validate:"required"`
	Amenities    []string `json:"amenities,omitempty"`
	Images       []string `json:"images,omitempty"`
}

type HotelsDto []HotelDto
