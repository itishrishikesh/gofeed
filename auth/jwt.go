package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

var Key string = "dadsgfjahg2131623879shdgrfahjdajhgfjafajhsdgfajhgdfjasghf"

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	log.Println("Time when JWT token was generated", time.Now().Add(10000*time.Minute).String())
	claims["exp"] = time.Now().Add(10000 * time.Minute)
	claims["authorized"] = true
	claims["user"] = username
	token.Header["exp"] = time.Now().Add(10000 * time.Minute)
	tokenString, err := token.SignedString([]byte(Key))
	if err != nil {
		log.Println("E#1PBTIY - Failed to Generate JWT Token", err)
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) *jwt.Token {
	log.Println("Time when JWT token was being verified", time.Now())
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
