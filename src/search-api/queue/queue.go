package queue

import (
	"bytes"
	"encoding/json"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"

	hotelService "hotel-api/service"
	"io/ioutil"
	"net/http"
	sdto "search-api/searchDto"
)

var queue amqp.Queue
var channel *amqp.Channel

func InitQueue() {

	// Configura la conexión a RabbitMQ y declara una cola.
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		log.Info("Failed to connect to RabbitMQ")
		log.Fatal(err)
	} else {
		log.Info("RabbitMQ connection established")
	}

	channel, err = conn.Channel() // Abre un canal de comunicación
	if err != nil {
		log.Info("Failed to open channel")
		log.Fatal(err)
	} else {
		log.Info("Channel opened")
	}

	queue, err = channel.QueueDeclare(
		"hotel",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Info("Failed to declare a queue")
		log.Fatal(err)
	} else {
		log.Info("Queue declared")
	}
}

// ProcesarMensaje procesa el mensaje y envía los datos a Solr.
func ProcesarMensaje(jsonMessage sdto.QueueMessageDto) {

	// Procesar el mensaje y obtener datos en formato Solr
	datosSolr := prepararDatosParaSolr(jsonMessage)

	// Enviar datos a Solr
	err := EnviarDatosASolr(datosSolr)
	if err != nil {
		log.Error("Error al enviar datos a Solr:", err)
	}
}

// prepararDatosParaSolr convierte el objeto HotelDto en un formato adecuado para Solr.
func prepararDatosParaSolr(messageDto sdto.QueueMessageDto) interface{} {
	hotelDto, err := hotelService.HotelService.GetHotelById(messageDto.Id)
	if err != nil {
		//TODO error
	}

	//TODO obtener un HotelDto desde la DB usando el ID

	// Creamos un mapa para representar un documento Solr
	documentoSolr := make(map[string]interface{})

	// Añadimos campos del objeto HotelDto al documento Solr
	documentoSolr["id"] = hotelDto.Id
	documentoSolr["nombre_hotel"] = hotelDto.Name
	documentoSolr["ciudad"] = hotelDto.City
	documentoSolr["descripcion"] = hotelDto.Description
	documentoSolr["cantidad_habitaciones"] = hotelDto.RoomAmount
	documentoSolr["calle"] = hotelDto.StreetName
	documentoSolr["numero_calle"] = hotelDto.StreetNumber
	documentoSolr["rate"] = hotelDto.Rate
	documentoSolr["amenities"] = hotelDto.Amenities
	documentoSolr["imagenes"] = hotelDto.Images

	return documentoSolr
}

// EnviarDatosASolr envía los datos a Solr.
func EnviarDatosASolr(datos interface{}) error {

	url := "http://localhost:8983/solr/mi_indice/update?commit=true" // Ajusta la URL según tu configuración

	// Convertir los datos a formato JSON
	jsonData, err := json.Marshal(datos)
	if err != nil {
		return err
	}

	// Realizar una solicitud HTTP POST a Solr
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Leer la respuesta de Solr
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Verificar si Solr respondió correctamente
	if resp.StatusCode != http.StatusOK {
		return errors.New("Solr request failed: " + string(body))
	}

	return nil
}

// Consume escucha mensajes de la cola y los procesa.
func Consume() {

	listaMensajes, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		log.Error("Error al consumir mensajes:", err)
	}

	for mensaje := range listaMensajes {
		var jsonMessage sdto.QueueMessageDto

		err = json.Unmarshal(mensaje.Body, &jsonMessage)
		if err != nil {
			log.Error("Error al decodificar mensaje:", err)
		}

		// Procesar el mensaje
		ProcesarMensaje(jsonMessage)
	}
}
