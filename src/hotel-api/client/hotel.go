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

type hotelClient struct{}

// estructura vacía que se utiliza para implementar el receptor de los métodos del cliente de hoteles.

// hotelClientInterface define la interfaz para el cliente de hoteles
type hotelClientInterface interface {
	InsertHotel(hotel model.Hotel) model.Hotel
	GetHotelById(id string) model.Hotel
	GetHotels() model.Hotels
	DeleteHotelById(id string) error
	UpdateHotelById(hotel model.Hotel) model.Hotel
}

var HotelClient hotelClientInterface

func init() {
	HotelClient = &hotelClient{}
}

// InsertHotel inserta un nuevo hotel en la base de datos MongoDB
func (c hotelClient) InsertHotel(hotel model.Hotel) model.Hotel {

	db := db.MongoDb
	insertHotel := hotel
	insertHotel.Id = primitive.NewObjectID()

	_, err := db.Collection("hotels").InsertOne(context.TODO(), &insertHotel)

	if err != nil {
		fmt.Println(err)
		return hotel
	}

	hotel.Id = insertHotel.Id
	return hotel
}

// GetHotelById obtiene un hotel por su ID desde la base de datos MongoDB
func (c hotelClient) GetHotelById(id string) model.Hotel {
	var hotel model.Hotel
	db := db.MongoDb
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println(err)
		return hotel
	}

	err = db.Collection("hotels").FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&hotel)
	if err != nil {
		fmt.Println(err)
		return hotel
	}
	return hotel
}

// GetHotels obtiene todos los hoteles desde la base de datos MongoDB
func (c hotelClient) GetHotels() model.Hotels {
	var hotels model.Hotels
	db := db.MongoDb
	cursor, err := db.Collection("hotels").Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Println(err)
		return hotels
	}

	err = cursor.All(context.TODO(), &hotels)

	if err != nil {
		fmt.Println(err)
		return hotels
	}
	return hotels
}

// DeleteHotelById elimina un hotel por su ID desde la base de datos MongoDB
func (c hotelClient) DeleteHotelById(id string) error {

	db := db.MongoDb
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := db.Collection("hotels").DeleteOne(context.TODO(), bson.D{{"_id", objID}})

	if result.DeletedCount == 0 {
		log.Debug("Hotel not found")
		return errors.New("hotel not found")

	} else if err != nil {
		log.Debug("Failed to delete hotel")

	} else {
		log.Debug("Hotel deleted successfully: ", id)
	}
	return err
}

// UpdateHotelById actualiza un hotel por su ID en la base de datos MongoDB
func (c hotelClient) UpdateHotelById(hotel model.Hotel) model.Hotel {

	db := db.MongoDb
	update := bson.D{{"$set",
		bson.D{
			{"name", hotel.Name},
			{"room_amount", hotel.RoomAmount},
			{"description", hotel.Description},
			{"city", hotel.City},
			{"street_name", hotel.StreetName},
			{"street_number", hotel.StreetNumber},
			{"rate", hotel.Rate},
			{"amenities", hotel.Amenities},
			{"images", hotel.Images},
		},
	}}

	result, err := db.Collection("hotels").UpdateOne(context.TODO(), bson.D{{"_id", hotel.Id}}, update)

	if result.MatchedCount != 0 {
		log.Debug("Updated hotel successfully")
		return hotel

	} else if err != nil {
		log.Debug("Failed to update hotel")
	}
	return model.Hotel{}
}
