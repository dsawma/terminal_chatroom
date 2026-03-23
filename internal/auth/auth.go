package auth

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/alexedwards/argon2id"
	"github.com/dsawma/terminal_chatroom/internal/database"
)

func Login(ctx context.Context, q *database.Queries) (string, error) {
	fmt.Println("Welcome")
	fmt.Println("Login or Signup")
	register_words := GetInput()
	if len(register_words) == 0 {
		return "", errors.New("you must enter an option. goodbye")
	}
	register := register_words[0]

	switch register {
	case "Login":
		fmt.Println("Please enter your username:")
		u_words := GetInput()
		if len(u_words) == 0 {
			return "", errors.New("you must enter a username. goodbye")
		}
		username := u_words[0]
		fmt.Println("Please enter your password:")
		p_words := GetInput()
		if len(p_words) == 0 {
			return "", errors.New("you must enter a password. goodbye")
		}
		password := p_words[0]

		user, err := q.GetUserByUsername(ctx, username)
		if err != nil {
			return "", errors.New("Incorrect Username or Password")
		}

		valid, err := CheckPasswordHash(password, user.HashedPassword)
		if err != nil || !valid {
			return "", errors.New("Incorrect Username or Password")
		}
		return user.Username, nil

	case "Signup":
		fmt.Println("Please enter your username:")
		u_words := GetInput()
		if len(u_words) == 0 {
			return "", errors.New("you must enter a username. goodbye")
		}
		username := u_words[0]
		fmt.Println("Please enter your password:")
		p_words := GetInput()
		if len(p_words) == 0 {
			return "", errors.New("you must enter a password. goodbye")
		}
		password := p_words[0]

		_, err := q.GetUserByUsername(ctx, username)
		if err == nil{
			return "", errors.New("User already exist")
			
		}

		hashed_password, err := HashPassword(password)
		if err != nil {
			return "", errors.New("Failed to create hashed_password")
		}

		user, err := q.CreateUser(ctx, database.CreateUserParams{Username: username, HashedPassword: hashed_password})
		if err != nil {
			fmt.Println("DB ERROR:", err) 
        	return "", err
		}
		return user.Username, nil
	default:
		return "", errors.New("Wrong input option")

	}

}
func HashPassword(password string) (string, error) {
	
	hash, err := argon2id.CreateHash(password, &argon2id.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2, 
		SaltLength:  16,
		KeyLength:   32,
	})
	if err != nil {
		log.Printf("DEBUG: Database error: %v", err)
		return "", err
	}
	return hash, nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return match, err
	}
	return match, nil
}

func GetInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		return nil
	}
	line := scanner.Text()
	line = strings.TrimSpace(line)
	return strings.Fields(line)

}


func JoinRoom(ctx context.Context, q *database.Queries) (string, error) {
	fmt.Println("CREATE new ChatRoom or JOIN existing?")
	response_word := GetInput()
	response := response_word[0]
	switch response{
	case "create":
		fmt.Println("What is the ChatRoom name?")
		create_word := GetInput()
		create := create_word[0]
		room, err := q.CreateRoom(ctx, create)
		if err == nil{
			return "", errors.New("Failed to create Room")
		} else{
		return room.RoomName, nil
		}
	case "join":
		lstRooms,err := q.GetAllRoomNames(ctx)
		if err == nil{
			return "", errors.New("Failed to fetch rooms")
		}
		fmt.Println("Available Rooms:")
		for i, room := range lstRooms{
			fmt.Printf("%d. %s\n", i + 1, room)
		}
		fmt.Println("JOIN or CREATE a new room")
		newResp_word := GetInput()
		newResp := newResp_word[0]
		switch newResp{
		case "join":
			fmt.Println("Type the Room Number to join")
			roomNum_word:= GetInput()
			roomNum := roomNum_word[0]
			num, err := strconv.Atoi(roomNum)
			if err != nil{
				return "", errors.New("Failed to get room num")
			}
			indexNum := num -1
			if indexNum >= 0 && indexNum <= len(lstRooms) -1{
				return lstRooms[indexNum], nil
			}else{
				return "", errors.New("Invalid range of Room Number")
			}
		case "create":
			fmt.Println("What is the ChatRoom name?")
			create_word := GetInput()
			create := create_word[0]
			room, err := q.CreateRoom(ctx, create)
			if err == nil{
				return "", errors.New("Failed to create Room")
			}else {
			return room.RoomName, nil
			}
		default:
			return "", errors.New("Wrong input option")
		}

	default:
		return "", errors.New("Wrong input option")
	}
}

