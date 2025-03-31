package core

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitMQConn *amqp.Connection
var RabbitMQChannel *amqp.Channel

func InitRabbitMQ() {
	var err error
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	RabbitMQConn, err = amqp.Dial(url)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	RabbitMQChannel, err = RabbitMQConn.Channel()
	if err != nil {
		log.Fatal("Failed to open channel:", err)
	}

	exchangeName := os.Getenv("RABBITMQ_EXCHANGE")
	// Declare exchange
	err = RabbitMQChannel.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare exchange:", err)
	}

	// Declare queues and bind them to the exchange
	queueKY026 := os.Getenv("RABBITMQ_QUEUE_KY026")
	queueMQ2 := os.Getenv("RABBITMQ_QUEUE_MQ2")
	queueMQ135 := os.Getenv("RABBITMQ_QUEUE_MQ135")
	queueDHT22 := os.Getenv("RABBITMQ_QUEUE_DHT22")

	// Declare KY026 queue
	_, err = RabbitMQChannel.QueueDeclare(
		queueKY026,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare KY026 queue:", err)
	}

	// Bind KY026 queue to exchange
	err = RabbitMQChannel.QueueBind(
		queueKY026,
		queueKY026, // Use queue name as routing key
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to bind KY026 queue:", err)
	}

	// Declare MQ2 queue
	_, err = RabbitMQChannel.QueueDeclare(
		queueMQ2,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare MQ2 queue:", err)
	}

	// Bind MQ2 queue to exchange
	err = RabbitMQChannel.QueueBind(
		queueMQ2,
		queueMQ2, // Use queue name as routing key
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to bind MQ2 queue:", err)
	}

	// Declare MQ135 queue
	_, err = RabbitMQChannel.QueueDeclare(
		queueMQ135,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare MQ135 queue:", err)
	}

	// Bind MQ135 queue to exchange
	err = RabbitMQChannel.QueueBind(
		queueMQ135,
		queueMQ135, // Use queue name as routing key
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to bind MQ135 queue:", err)
	}

	// Declare DHT22 queue
	_, err = RabbitMQChannel.QueueDeclare(
		queueDHT22,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare DHT22 queue:", err)
	}

	// Bind DHT22 queue to exchange
	err = RabbitMQChannel.QueueBind(
		queueDHT22,
		queueDHT22, // Use queue name as routing key
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to bind DHT22 queue:", err)
	}

	log.Printf("Successfully connected to RabbitMQ and set up queues: %s, %s, %s, %s", queueKY026, queueMQ2, queueMQ135, queueDHT22)
}

func PublishMessage(queue string, message []byte) error {
	log.Printf("Publishing message to queue %s: %s", queue, string(message))

	err := RabbitMQChannel.Publish(
		os.Getenv("RABBITMQ_EXCHANGE"),
		queue, // Use queue name as routing key
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         message,
			DeliveryMode: amqp.Persistent,
		},
	)

	if err != nil {
		log.Printf("Error publishing message: %v", err)
		return err
	}

	log.Printf("Successfully published message to queue %s", queue)
	return nil
}
