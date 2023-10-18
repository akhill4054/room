package util

import (
	"fmt"
	"time"

	"github.com/akhill4054/room-backend/config"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	Uid      int    `json:"uid"`
	Username string `json:"username"`
}

func GenerateToken(uid int, username string) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour * 24 * 30)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
		Uid:      uid,
		Username: username,
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenObj.SignedString([]byte(config.JWT_SECRET))

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims := &Claims{}

	_, err := jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method")
		}
		return []byte(config.JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	err = tokenClaims.RegisteredClaims.Valid()
	return tokenClaims, nil
}
