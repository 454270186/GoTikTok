package rabbitmq

import (
	"context"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

// User direct exchange (default)
// Every RabbitMQ instance has unique queue name
type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	QueueName string
}

// Create a RabbitMQ instance in Simple mode
func NewSimple() *RabbitMQ {
	rabbitMQ := &RabbitMQ{}
	var err error
	rabbitMQ.conn, err = amqp.Dial(mqAddr)
	failOnError(err, "Fail to connect to RabbitMQ")

	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	failOnError(err, "Fail to open a channel")

	rabbitMQ.QueueName = queName

	return rabbitMQ
}

// Publish msg to default exchange using queue name for routing
func (r *RabbitMQ) PubWithCtx(ctx context.Context, body []byte) error {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return errors.New("fail to declare queue: " + err.Error())
	}

	err = r.channel.PublishWithContext(ctx, 
		"", // user default exchange
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: body,
		},
	)
	if err != nil {
		return errors.New("fail to publish a msg: " + err.Error())
	}

	return nil
}

//  Get msgs from default exchange using queue name for routing
func (r *RabbitMQ) GetMsgs() (<- chan amqp.Delivery, error) {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.New("fail to declare queue: " + err.Error())
	}

	msgs, err := r.channel.Consume(
		r.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.New("fail to get msgs from queue: " + err.Error())
	}

	return msgs, nil
}
