package app

import (
	"auth-service/internal/messaging"
	"auth-service/internal/messaging/consumer"
	"log"
)

func RegisterConsumers(container *Container) {

	if err := messaging.DeclareExchanges(container.Rabbit.Channel); err != nil {
		log.Fatal(err)
	}
	
	consumer.ConsumeUserUpdated(
		container.Rabbit.Channel,
		container.AuthService,
	)
}
