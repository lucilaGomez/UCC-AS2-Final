package client

import (
	"context"
	"errors"
	"fmt"
	db "hotel-api/db"
	"hotel-api/model"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type amenityClient struct{}

// amenityClientInterface define la interfaz para el cliente de amenidades
type amenityClientInterface interface {
	InsertAmenity(amenity model.Amenity) model.Amenity
	GetAmenityById(id string) model.Amenity
	GetAmenityByName(name string) model.Amenity
	GetAmenities() model.Amenities
	DeleteAmenityById(id string) error
	//UpdateHotel(hotel model.Hotel) model.Hotel
}

var AmenityClient amenityClientInterface

// Define una interfaz para el cliente de amenidades

// estructura vacía que se utiliza para implementar el receptor de los métodos del cliente de amenidades.
func init() {
	AmenityClient = &amenityClient{}
}

// InsertAmenity inserta una amenidad en la base de datos MongoDB
func (c amenityClient) InsertAmenity(amenity model.Amenity) model.Amenity {

	db := db.MongoDb

	insertAmenity := amenity
	insertAmenity.Id = primitive.NewObjectID()

	_, err := db.Collection("amenities").InsertOne(context.TODO(), &insertAmenity)

	if err != nil {
		fmt.Println(err)
		return amenity
	}

	amenity.Id = insertAmenity.Id

	return amenity
}

// GetAmenityById obtiene una amenidad por su ID desde la base de datos MongoDB
func (c amenityClient) GetAmenityById(id string) model.Amenity {
	var amenity model.Amenity

	db := db.MongoDb
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println(err)
		return amenity
	}

	err = db.Collection("amenities").FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&amenity)
	if err != nil {
		fmt.Println(err)
		return amenity
	}
	return amenity
}

// GetAmenityByName obtiene una amenidad por su nombre desde la base de datos MongoDB
func (c amenityClient) GetAmenityByName(name string) model.Amenity {
	var amenity model.Amenity

	db := db.MongoDb

	err := db.Collection("amenities").FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&amenity)

	if err != nil {
		fmt.Println(err)
		return amenity
	}
	return amenity
}

// GetAmenities obtiene todas las amenidades desde la base de datos MongoDB
func (c amenityClient) GetAmenities() model.Amenities {
	var amenities model.Amenities

	db := db.MongoDb

	cursor, err := db.Collection("amenities").Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Println(err)
		return amenities
	}

	err = cursor.All(context.TODO(), &amenities)

	if err != nil {
		fmt.Println(err)
		return amenities
	}

	return amenities
}

// DeleteAmenityById elimina una amenidad por su ID desde la base de datos MongoDB
func (c amenityClient) DeleteAmenityById(id string) error {

	db := db.MongoDb

	objID, err := primitive.ObjectIDFromHex(id)

	result, err := db.Collection("amenities").DeleteOne(context.TODO(), bson.D{{"_id", objID}})

	if result.DeletedCount == 0 {
		log.Debug("Amenity not found")
		return errors.New("amenity not found")

	} else if err != nil {
		log.Debug("Failed to delete amenity")

	} else {
		log.Debug("Amenity deleted successfully: ", id)
	}
	return err
}
