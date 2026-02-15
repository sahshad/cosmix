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
			return conn
		}
		log.Printf("RabbitMQ connection failed (attempt %d/15): %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	log.Fatal("Failed to connect to RabbitMQ after multiple attempts:", err)
	return nil
}

func NewRabbitMQChannel(url string) *amqp.Channel {
	conn := ConnectRabbitMQ(url)
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("RabbitMQ channel failed:", err)
	}

	return ch
}
