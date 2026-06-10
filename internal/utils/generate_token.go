package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Jwtkey []byte

func Init(key string) {
	Jwtkey = []byte(key)
}
func GenerateToken(id int) (string, error) {

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Jwtkey)
}
