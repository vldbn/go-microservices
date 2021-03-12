package broker

import "github.com/streadway/amqp"

func SendMessage(channel *amqp.Channel, routingKey string, message []byte) error {
	return channel.Publish(
		"users",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
}
