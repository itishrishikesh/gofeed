package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var Key string = "dadsgfjahg2131623879shdgrfahjdajhgfjafajhsdgfajhgdfjasghf"

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = username
	tokenString, err := token.SignedString([]byte(Key))
	if err != nil {
		fmt.Println("E#1PBTIY - Failed to Generate JWT Token")
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
		fmt.Println("E#1PBU9G - Failed to verify token.")
		return nil
	}
	return token
}
