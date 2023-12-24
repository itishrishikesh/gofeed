package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

var Key string = "dadsgfjahg2131623879shdgrfahjdajhgfjafajhsdgfajhgdfjasghf"

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"user":       username,
		"exp":        time.Now().Add(10 * time.Minute).Unix(),
	})
	tokenString, err := token.SignedString([]byte(Key))
	if err != nil {
		log.Println("E#1PBTIY - Failed to Generate JWT Token", err)
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) *jwt.Token {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Key), nil
	})
	if err != nil {
		log.Println("E#1PBU9G - Failed to verify token.", err)
		return &jwt.Token{Valid: false}
	}
	return token
}
