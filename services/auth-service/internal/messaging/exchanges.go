package messaging

import amqp "github.com/rabbitmq/amqp091-go"

func DeclareExchanges(ch *amqp.Channel) error {

	// Auth owns this exchange
	if err := ch.ExchangeDeclare(
		"auth.events",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	// Optional DLX
	if err := ch.ExchangeDeclare(
		"auth.events.dlx",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}
