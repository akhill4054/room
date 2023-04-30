package util

import (
	"errors"

	"github.com/akhill4054/room-backend/config"
	"golang.org/x/crypto/bcrypt"
)

func GetPasswordHash(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(config.PASSWORD_SALT+password), bcrypt.DefaultCost,
	)
	return hashedPassword, err
}

func ComparePasswords(plainPassword string, hashedPassword string) bool {
	password := config.PASSWORD_SALT + plainPassword
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(password),
	)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false
		}
		panic(err)
	}

	return true
}
