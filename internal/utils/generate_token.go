package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var Jwtkey []byte

func Init(key string) {
	Jwtkey = []byte(key)
}

func GenerateToken(id int) (string, string, error) {

	sessionId := uuid.New().String()

	claims := jwt.MapClaims{
		"id":         id,
		"session_id": sessionId,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(Jwtkey)
	if err != nil {
		return "", "", err
	}

	return signedToken, sessionId, nil
}
