package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// ValidateToken func gets tokenString, secretKey and id, parses it
// then it checks if it valid or not and return parsed token pointer
// and bool (true == valid/false == invalid) and error if occurs.
func ValidateToken(tokenString, secretKey string, id int64) (*jwt.Token, bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, false, fmt.Errorf("error validating token: %s", err)
	}

	// When we parse token parser check if token is valid
	if !token.Valid {
		return nil, false, errors.New("invalid token")
	}



	return token, true, nil
}