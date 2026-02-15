package events

import (
	"encoding/json"
	"log"

	"user-service/internal/dto"
	"user-service/internal/services"

	authEvents "cosmix-events/auth"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeUserCreated(ch *amqp.Channel, userProfileService services.UserProfileService) {
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
			var event authEvents.UserCreated
			json.Unmarshal(msg.Body, &event)

			userCreatedEvent := dto.UserCreatedFromDTO{
				AuthUserID: event.AuthUserID,
				Email:      event.Email,
				FirstName:  event.FirstName,
				LastName:   event.LastName,
				CreatedAt:  event.CreatedAt,
			}

			err := userProfileService.CreateFromAuthEvent(userCreatedEvent)

			if err != nil {
				log.Println("User creation failed:", err)
			}
		}
	}()
}
