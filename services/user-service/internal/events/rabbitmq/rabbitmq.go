package events

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQ(url string) *amqp.Connection {
	var conn *amqp.Connection
	var err error

	for i := range 15 {
		conn, err = amqp.Dial(url)
		if err == nil {
			log.Println("Successfully connected to RabbitMQ")
			return conn
		}
		log.Printf("RabbitMQ connection failed (attempt %d/15): %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	log.Printf("Failed to connect to RabbitMQ after multiple attempts: %v", err)
	return nil
}

func NewRabbitMQChannel(url string) *amqp.Channel {
	conn := ConnectRabbitMQ(url)
	if conn == nil {
		return nil
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("RabbitMQ channel failed: %v", err)
		return nil
	}

	_, err = ch.QueueDeclare(
		"auth.user.created",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("Queue declare failed: %v", err)
		return nil
	}

	return ch
}
