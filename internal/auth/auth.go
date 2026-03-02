package auth

import (
	"bufio"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/alexedwards/argon2id"
	"github.com/dsawma/Sentinel-/internal/database"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func Login(ctx context.Context, q *database.Queries) (string, error) {
	fmt.Println("Welcome")
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
    user, err := q.GetUserByUsername(ctx, username)
    if err != nil {
        return "", errors.New("Incorrect Username or Password")
    }
	valid, err := CheckPasswordHash(password, user.HashedPassword)
	if err != nil || !valid {
    	return "", errors.New("Incorrect Username or Password")
	}
	token, err := MakeJWT( user.ID, cfg.jwtSecret, time.Duration(3600)*time.Second)
	if err != nil {
		return "", errors.New("Can't create Token")
	}
	
	_, err = ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
    	return "", errors.New("Invalid Token")
	}
	_, err = MakeRefreshToken() 
	if err != nil {
    	return "", errors.New("Couldn't make refresh")
	}

	return user.Username, nil

}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error){
	now := time.Now().UTC()
	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(expiresIn)),
		Subject: userID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken,err := token.SignedString([]byte(tokenSecret))
	if err != nil{
		return "", err
	}
	return signedToken, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error){
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
	return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	} 
	claim,ok := token.Claims.(*jwt.RegisteredClaims) 
	if !ok || !token.Valid{
		return uuid.Nil, errors.New("invalid token")
	}
	retUID,err:= uuid.Parse(claim.Subject)
	if err != nil{
		return uuid.Nil, err
	}
	return retUID, nil
}


func MakeRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	hex := hex.EncodeToString(bytes)
	return hex, nil
}


func CheckPasswordHash(password, hash string)(bool, error){
	match, err:= argon2id.ComparePasswordAndHash(password, hash) 
	if err!= nil{
		return match, err
	}
	return match, nil
}

func GetInput() []string {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned{
		return nil 
	}
	line := scanner.Text() 
	line = strings.TrimSpace(line) 
	return strings.Fields(line)
	
}