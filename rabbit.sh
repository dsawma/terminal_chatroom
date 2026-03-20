#!/bin/bash

start_or_run () {
    docker inspect chat_rabbitmq > /dev/null 2>&1

    if [ $? -eq 0 ]; then
        echo "Starting Chat RabbitMQ container..."
        docker start chat_rabbitmq
    else
        echo "Chat RabbitMQ container not found, creating a new one..."
        docker run -d --name chat_rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.13-management
    fi
}

case "$1" in
    start)
        start_or_run
        ;;
    stop)
        echo "Stopping Chat RabbitMQ container..."
        docker stop chat_rabbitmq
        ;;
    logs)
        echo "Fetching logs for Chat RabbitMQ container..."
        docker logs -f chat_rabbitmq
        ;;
    *)
        echo "Usage: $0 {start|stop|logs}"
        exit 1
esac
