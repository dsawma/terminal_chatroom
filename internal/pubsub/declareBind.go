package pubsub

import (

	amqp "github.com/rabbitmq/amqp091-go"
)
type SimpleQueueType int 
const (
	DurableQueue SimpleQueueType = iota
	TransientQueue 
)


func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType,
) (*amqp.Channel, amqp.Queue, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}
	queue := amqp.Queue{}
	switch queueType{
		case DurableQueue: 
			queue, err = channel.QueueDeclare(
			queueName, true, false, false, false, amqp.Table{"x-dead-letter-exchange": "chat_dlx"})
			if err != nil {
				return nil,amqp.Queue{} , err
			}

		case TransientQueue:
			queue, err = channel.QueueDeclare(
			queueName, false, true, true, false, amqp.Table{"x-dead-letter-exchange": "chat_dlx"})
			if err != nil {
				return nil, amqp.Queue{}, err
			}
	}
	err = channel.QueueBind(queueName, key, exchange, false, nil )
	if err != nil {
		return nil, amqp.Queue{}, err
	}
	return channel, queue, nil 
}