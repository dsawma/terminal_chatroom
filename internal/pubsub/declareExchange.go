package pubsub

import (

	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareExchange(ch *amqp.Channel, exchangeName string) error {
    return ch.ExchangeDeclare(
        exchangeName, // name
        "topic",      // type (match this to your routing!)
        true,         // durable
        false,        // auto-deleted
        false,        // internal
        false,        // no-wait
        nil,          // arguments
    )
}