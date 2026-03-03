package auth

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
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
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
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
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		return nil
	}
	line := scanner.Text()
	line = strings.TrimSpace(line)
	return strings.Fields(line)

}
