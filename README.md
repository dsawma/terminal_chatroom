# Terminal Chatroom 
This is a real time command-line chatroom that can connect with multiple clients. Using RabbitMQ, a pub-sub architecture is implemented so only users in a specific chatroom can receive messages. Also with PostreSQL, a user's account and created chatrooms are saved after terminating the program. 

# Motivation 
After learning RabbitMQ and PosterSQL, I wanted to create something that saved prexisting data while using a distributive system such as a pub-sub architecture. I thought creating a chatroom application would be a fun and interesting way to tackle this project. 

# Quick Start 

# Usage
When inside the chatroom
- '/q' - to terminate the program
- '/join [chatroom]' - to join another chatroom(assuming chatroom exists)
- Type anything to send message

# Contributing
git clone https://github.com/dsawma/terminal_chatroom
## Start rabbitMQ with Docker
./rabbitmq.sh
## build the client and server 
go build ./cmd/client/main.go
go build ./cmd/server/main.go
## Run the client(s) and server
go run ./cmd/client/.
go run ./cmd/server/.

