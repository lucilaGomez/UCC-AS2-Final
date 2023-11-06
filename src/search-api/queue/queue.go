package queue

import (
	"bytes"
	"encoding/json"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"search-api/solr"

	sdto "search-api/searchDto"
)

var queue amqp.Queue
var channel *amqp.Channel

func InitQueue() {

	// Configura la conexión a RabbitMQ y declara una cola.
	conn, err := amqp.Dial("amqp://user:password@rabbitmq:5672/")

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

	solr.SolrClient.Add(jsonMessage)

	// Procesar el mensaje y obtener datos en formato Solr
	datosSolr := prepararDatosParaSolr(jsonMessage)

	// Enviar datos a Solr
	err := EnviarDatosASolr(datosSolr)
	if err != nil {
		log.Error("Error al enviar datos a Solr:", err)
	}
}

// prepararDatosParaSolr convierte el objeto HotelDto en un formato adecuado para Solr.
func prepararDatosParaSolr(messageDto sdto.QueueMessageDto) map[string]interface{} {

	// Creamos un mapa para representar un documento Solr
	documentoSolr := make(map[string]interface{})

	documentoSolr["id"] = messageDto.Id
	documentoSolr["name"] = messageDto.Name
	documentoSolr["city"] = messageDto.City
	documentoSolr["description"] = messageDto.Description
	documentoSolr["room_amount"] = messageDto.RoomAmount
	documentoSolr["street_name"] = messageDto.StreetName
	documentoSolr["street_number"] = messageDto.StreetNumber
	documentoSolr["rate"] = messageDto.Rate

	//TODO ver si hace falta amenities en solr (images no va a solr)
	return documentoSolr
}

// EnviarDatosASolr envía los datos a Solr.
func EnviarDatosASolr(docSolr interface{}) error {
	url := "http://localhost:8983/solr/hotels/update/json/docs?commit=true"

	//Convertir los docSolr a formato JSON
	jsonData, err := json.Marshal(docSolr)
	if err != nil {
		log.Error("Error al convertir mensaje a JSON:", err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Error("Error al enviar request post a Solr:", err)
		return err
	}
	defer resp.Body.Close()

	// Leer la respuesta de Solr
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error al leer body de respuesta de Solr:", err)
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
			log.Error("Error al decodificar mensaje de la cola:", err)
		}

		ProcesarMensaje(jsonMessage)
	}

}
