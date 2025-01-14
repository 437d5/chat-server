package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	ID int64
	Username string
	jwt.RegisteredClaims
}

// CreateToken func creates new JWT token with provided id, username and
// expiration time. It also needs secretKey from config to be provided.
// This function return JWT token of string type and error if occurs.
func CreateToken(id int64, username, secretKey string, expAt int) (string, error) {
	now := time.Now()
	exp := now.Add(time.Duration(expAt) * time.Hour)

	claims := TokenClaims{
		ID: id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed create new token: %s", err)
	}

	return tokenString, nil
}