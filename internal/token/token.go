package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type authClaims struct {
	UserId int `json:"id"`
	jwt.StandardClaims
}

func CreateToken(userId int, key string) (string, error) {
	claims := authClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 3).Unix(),
		},
	}
	oleg := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := oleg.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetClaims(token string, key string) (*authClaims, error) {
	var claims authClaims

	tokenParse, err := jwt.ParseWithClaims(
		token,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
	if err != nil {
		return nil, err
	}
	if !tokenParse.Valid {
		return nil, err
	}

	return &claims, nil
}
