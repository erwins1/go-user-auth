package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password, salt string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[auth.HashPassword] err hashing, err : %v", err)
		return string(result), err
	}
	return string(result), nil
}

func ValidatePassword(password, userPassword, salt string) error {
	saltedPassword := password + salt
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(saltedPassword))
	if err != nil {
		return err
	}
	return nil
}
