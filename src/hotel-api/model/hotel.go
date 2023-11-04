package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	Id           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	RoomAmount   int                `bson:"room_amount"`
	Description  string             `bson:"description"`
	City         string             `bson:"city"`
	StreetName   string             `bson:"street_name"`
	StreetNumber int                `bson:"street_number"`
	Rate         float64            `bson:"rate"`
	Amenities    Amenities          `bson:"amenities"`
	Images       []string           `bson:"images"`
}

type Hotels []Hotel
