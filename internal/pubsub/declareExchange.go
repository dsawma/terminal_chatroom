package pubsub

import (

	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareExchange(ch *amqp.Channel,name, exchangeType string) error {
    return ch.ExchangeDeclare(
        name,
        exchangeType,   
        true,         // durable
        false,        // auto-deleted
        false,        // internal
        false,        // no-wait
        nil,          // arguments
    )
}