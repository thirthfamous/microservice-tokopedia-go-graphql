package helper

import (
	"log"
	"time"

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

func GenerateToken(profileId int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["profile_id"] = profileId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}
