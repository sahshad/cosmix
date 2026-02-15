package publisher

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	authEvents "cosmix-events/auth"
)

func PublishUserUpdated(ch *amqp.Channel, event authEvents.UserUpdated) {
	body, _ := json.Marshal(event)

	err := ch.PublishWithContext(
		context.Background(),
		"auth.events",
		"user.updated",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println("Failed to publish user.updated event:", err)
	}
}
