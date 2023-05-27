package helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateNewJWT(iss string, sub string, key string, exp time.Duration) string {

	claims := jwt.MapClaims{
		"iss": iss,
		"sub": sub,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(exp).Unix(),
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the JWT token with a secret key
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		fmt.Println(err)
	}

	return signedToken
}

func ValidateJWT(tokenString string, key string) bool {
	// Verify the token with a secret ke
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(key), nil
	})

	if err != nil {
		fmt.Println(err)
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return true

	} else {
		return false
	}
}
