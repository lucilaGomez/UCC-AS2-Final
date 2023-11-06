package queue

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

var queue amqp.Queue
var channel *amqp.Channel

func InitQueue() {
	// Inicializa la conexión con el servidor RabbitMQ y declara una cola

	conn, err := amqp.Dial("amqp://user:password@rabbitmq:5672/")

	if err != nil {
		log.Info("Failed to connect to RabbitMQ")
		log.Fatal(err)
	} else {
		log.Info("RabbitMQ connection established")
	}

	channel, err = conn.Channel()
	// para abrir un canal de comunicación con RabbitMQ sobre la conexión establecida.

	if err != nil {
		log.Info("Failed to open channel")
		log.Fatal(err)
	} else {
		log.Info("Channel opened")
	}

	queue, err = channel.QueueDeclare(
		// Para declarar una cola en RabbitMQ.
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

func Publish(body []byte) error {
	// Para publicar un mensaje en la cola RabbitMQ

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// Establecer un límite de tiempo de 5 segundos para la operación de publicación.
	defer cancel()

	err := channel.PublishWithContext(
		// Publicar un mensaje en la cola especificada
		ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		log.Debug("Error while publishing message", err)
		return err
	}

	return nil
}
