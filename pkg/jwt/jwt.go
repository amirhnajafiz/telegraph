package jwt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtKey = []byte("my_secret_key")

type JWT struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(user string) (string, error) {
	// create token
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &JWT{
		Username: user,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	return token.SignedString(jwtKey)
}

func ParseToken(jwtToken string) (bool, error) {
	// Initialize a new instance of `Claims`
	claims := &JWT{}
	tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return tkn.Valid, err
}
