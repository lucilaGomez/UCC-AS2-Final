package db

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Maneja la conexión y desconexión a una base de datos MongoDB
var MongoDb *mongo.Database
var client *mongo.Client

func Disconect_db() {
	// Desconecta el cliente de MongoDB
	client.Disconnect(context.TODO())
}

func Init_db() {

	// Inicializa la conexión a la base de datos MongoDB
	clientOpts := options.Client().ApplyURI("mongodb://root:pass@mongodatabase:27017/?authSource=admin&authMechanism=SCRAM-SHA-256")
	cli, err := mongo.Connect(context.TODO(), clientOpts)
	// Para establecer la conexión con la base de datos MongoDB
	client = cli

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	dbNames, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	// Se utiliza el cliente para obtener una lista de nombres de bases de datos disponibles

	if err != nil {
		log.Info("Failed to get databases available")
		log.Fatal(err)
	}

	MongoDb = client.Database("test")

	fmt.Println("Available datatabases:")
	fmt.Println(dbNames)

}
