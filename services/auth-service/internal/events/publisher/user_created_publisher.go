package events

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	authEvents "cosmix-events/auth"
)

func PublishUserCreated(ch *amqp.Channel, event authEvents.UserCreated) {
	body, _ := json.Marshal(event)

	err := ch.PublishWithContext(
		context.Background(),
		"",
		"auth.user.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println("Failed to publish user.created event:", err)
	}
}
