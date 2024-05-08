package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

func CreateToken(r *http.Request, id int64, isAccess bool) (string, error) {
	var expiration time.Duration
	if isAccess {
		expiration = Tokens(r).AccessExp
	} else {
		expiration = Tokens(r).RefreshExp
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": expiration,
		})

	return token.SignedString([]byte(Tokens(r).SecretKey))
}
func VerifyToken(r *http.Request, tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Tokens(r).SecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
