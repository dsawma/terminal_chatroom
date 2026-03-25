# Terminal Chatroom 
This is a real time terminal chatroom app that can connect with multiple users. 

## Motivation 
After learning RabbitMQ and SQL, I wanted to create something that saved prexisting data while using a distributed system such as a pub-sub architecture. I thought creating a chatroom application would be a fun and interesting way to tackle this project. 

## Goals
* Real-Time Messaging: Leverages Goroutines and RabbitMQ to enable asynchronous, low-latency messaging across multiple concurrent users.
* Distributed System: Utilizes a Pub-Sub model via RabbitMQ to route messages efficiently to designated chatrooms using topic exchanges.
* Streamlined Developer Experience: Employs Docker to containerize the infrastructure(RabbitMQ), ensuring a consistent setup with a simple CLI interface.
* Persistent Data Storage: Ensures long-term storage of user profiles and rooms using a PostgreSQL relational database.
* Thread-Safe State Management: Implements Mutexes (sync.RWMutex) to safely manage local application state across multiple background message handlers.

## Usage
### client
1. Login or Sign-up for a user
2. Create or join a pre-existing chatroom
3. When inside the chatroom
    - '/q' - to terminate the program
    - '/join [chatroom]' - to join another chatroom
    - Type anything to send message

### server
When connected to rabbitMQ
- 'pause' - to stop all users from sending messages
- 'resume' - to undo the pause command
- 'quit' - to terminate the program 

## Getting Started
### Prerequisities
- Docker Desktop
- Go 1.22+ 

git clone https://github.com/dsawma/terminal_chatroom

### Create a .env file(Need to create own Postgres DB_URL)
cp .env.example .env 

### Start rabbitMQ with Docker 
./rabbitmq.sh
### Create Postgres Database
./migrate.sh
### build the client and server 
go build ./cmd/client/main.go
go build ./cmd/server/main.go
### Run the client(s) and server
go run ./cmd/client/.
go run ./cmd/server/.

## Contributions
Contributions are welcome! If you have a suggestion that would make this project better, please fork the repo and create a pull request. 