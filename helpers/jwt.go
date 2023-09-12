package helpers

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string

func GenerateJwt(userName string) (string, error) {
	if len(jwtSecret) == 0 {
		jwtSecret = os.Getenv("JWT_SECRET")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"userName": userName,
	})
	return token.SignedString(jwtSecret)
}

func ParseJwt(tokenString string) (string, error) {
	if len(jwtSecret) == 0 {
		jwtSecret = os.Getenv("JWT_SECRET")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecret, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if err == nil && ok && token.Valid {
		return claims["userName"].(string), nil
	} else {
		return "", errors.New("JWT not valid")
	}
}
