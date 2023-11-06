package searchDto

type QueueMessageDto struct {
	Id           string   `json:"id"`
	Message      string   `json:"message"`
	Name         string   `json:"name"`
	RoomAmount   int      `json:"room_amount"`
	Description  string   `json:"description"`
	City         string   `json:"city"`
	StreetName   string   `json:"street_name"`
	StreetNumber int      `json:"street_number"`
	Rate         float64  `json:"rate"`
	Amenities    []string `json:"amenities,omitempty"`
	Images       []string `json:"images,omitempty"`
}

type QueueMessagesDto []QueueMessageDto
