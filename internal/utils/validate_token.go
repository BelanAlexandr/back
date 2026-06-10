package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return Jwtkey, nil
		})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
