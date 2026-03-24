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

    docker inspect chat_db > /dev/null 2>&1
    if [ $? -eq 0 ]; then 
        echo "Starting chat postgres..."
        docker start chat_db
    else
        echo "Creating chat postgres..."
        docker run -d --name chat_db -e POSTGRES_USER=chat_user -p 5432:5432 -e POSTGRES_PASSWORD=chat_pass -e POSTGRES_DB=terminal_chatroom postgres:16-alpine
    fi
}

case "$1" in
    start)
        start_or_run
        ;;
    stop)
        echo "Stopping Chat RabbitMQ container..."
        docker stop chat_rabbitmq chat_db
        ;;
    logs)
        echo "Fetching logs"
        docker logs -f chat_rabbitmq & docker logs -f chat_db
        ;;
    *)
        echo "Usage: $0 {start|stop|logs}"
        exit 1
esac
