package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwt_secret = os.Getenv("JWT_SECRET")

var jwtKey = []byte(jwt_secret)

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(userId string, username string) string {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UserID:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return ""
	}
	return tokenString

}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}
