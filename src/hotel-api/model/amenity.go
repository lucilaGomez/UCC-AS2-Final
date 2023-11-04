package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Amenity struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

type Amenities []Amenity
