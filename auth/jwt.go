package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"
	tokenString, err := token.SignedString([]byte(os.Getenv("KEY")))
	if err != nil {
		fmt.Println("E#1PBTIY - Failed to Generate JWT Token")
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string) *jwt.Token {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("KEY")), nil
	})
	if err != nil {
		fmt.Println("E#1PBU9G - Failed to verify token.")
	}
	return token
}
