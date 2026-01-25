package events

import (
	"encoding/json"
	"log"

	"user-service/internal/services"

	amqp "github.com/rabbitmq/amqp091-go"
)

type UserCreatedEvent struct {
	AuthUserID uint   `json:"auth_user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
}

func ConsumeUserCreated(ch *amqp.Channel, userProfileService services.UserProfileService) {
	log.Println("Consuming user.created events...")
	msgs, err := ch.Consume(
		"auth.user.created",
		"",
		true, // auto-ack (OK for MVP)
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("Consumer failed:", err)
	}

	go func() {
		for msg := range msgs {
			var event UserCreatedEvent
			json.Unmarshal(msg.Body, &event)

			err := userProfileService.CreateFromAuthEvent(
				event.AuthUserID,
				event.FirstName,
				event.LastName,
			)

			if err != nil {
				log.Println("User creation failed:", err)
			}
		}
	}()
}
