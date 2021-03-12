package broker

import "github.com/streadway/amqp"

// QueueDeclare declare queues
func QueueDeclare(ch *amqp.Channel, queues map[string]string) error {
	for q, _ := range queues {
		_, err := ch.QueueDeclare(
			q,
			false, // durable
			false, // delete when unused
			true,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// QueueBind bind route to queue
func QueueBind(ch *amqp.Channel, queues map[string]string) error {
	for q, r := range queues {
		err := ch.QueueBind(
			q,
			r,
			"users",
			false,
			nil,
		)

		if err != nil {
			return err
		}
	}
	return nil
}

// MakeMessageChannel create message channel
func MakeMessageChannel(ch *amqp.Channel, queue string) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		queue, // queue
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}
