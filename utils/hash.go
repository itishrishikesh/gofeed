package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil || hash == nil {
		log.Println("E#1PHA7U - Failed to hash the given string", err)
		return ""
	}
	return string(hash)
}

func CompareHashAndPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println("E#1PHAZP - Enable to hash", err)
		return false
	}
	return true
}
