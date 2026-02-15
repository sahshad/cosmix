package consumer

import (
	"encoding/json"
	"log"

	"auth-service/internal/dto"
	"auth-service/internal/services"

	authEvents "cosmix-events/auth"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeUserUpdated(ch *amqp.Channel, userService services.AuthServiceInterface) {
	q, err := ch.QueueDeclare(
		"user.updated",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Queue declaration failed:", err)
	}

	if err := ch.QueueBind(
		q.Name,
		"user.updated",
		"auth.events",
		false,
		nil,
	); err != nil {
		log.Fatal("Queue binding failed:", err)
	}

	msgs, err := ch.Consume(
		q.Name,
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
			var event authEvents.UserUpdated
			json.Unmarshal(msg.Body, &event)

			userUpdatedEvent := dto.UserUpdatedFromDTO{
				AuthUserID: event.AuthUserID,
				Email:      event.Email,
				UpdatedAt:  event.UpdatedAt,
			}

			err := userService.UpdateFromAuthEvent(userUpdatedEvent)

			if err != nil {
				log.Println("User update failed:", err)
			}
		}
	}()
}
