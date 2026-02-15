package app

import (
	"user-service/internal/messaging"
	"user-service/internal/messaging/consumer"
	"log"
)

func RegisterConsumers(container *Container) {

	if err := messaging.DeclareExchanges(container.Rabbit.Channel); err != nil {
		log.Fatal(err)
	}
	
	consumer.ConsumeUserCreated(
		container.Rabbit.Channel,
		container.UserProfileService,
	)
}