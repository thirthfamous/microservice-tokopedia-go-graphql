package helper

import (
	"github.com/golang-jwt/jwt"
)

var (
	SecretKey = []byte("secret")
)

func ParseToken(tokenStr string) (bool, float64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		profile_id := claims["profile_id"].(float64)
		return token.Valid, profile_id, nil
	} else {
		return false, 0, err
	}
}
