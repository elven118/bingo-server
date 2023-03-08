package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTManger struct {
	secretKey     string
	tokenDuration time.Duration
}

// custom claims
type UserClaims struct {
	jwt.RegisteredClaims
	Id   string `json:"id"`
	Name string `json:"name"`
}

// jwt secret key
var jwtSecret = []byte("secret")

func GenToken(id string, name string) (string, error) {
	// set claims and sign
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(300 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   name,
		},
		Id:   id,
		Name: name,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		log.Println("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (*UserClaims, error) {
	token, _ := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Println("Invalid Token")
		return nil, fmt.Errorf("Invalid Token")
	}
}

func CreateJWTManger(secretKey string, tokenDuration time.Duration) *JWTManger {
	return &JWTManger{secretKey, tokenDuration}
}
