package token

import (
	"fmt"
	"github.com/and67o/go-comments/internal/configuration"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type authClaims struct {
	UserId int `json:"id"`
	jwt.StandardClaims
}
type typeToken int

const (
	refreshType typeToken = iota
	accessType
)

func GetAccessKey(userId int) string {
	return fmt.Sprintf("%d-%d", userId, accessType)
}

func GetRefreshKey(userId int) string {
	return fmt.Sprintf("%d-%d", userId, refreshType)
}

type Tokens struct {
	AccessToken         string        `json:"access_token"`
	RefreshToken        string        `json:"refresh_token"`
	AccessTokenExpires  time.Duration `json:"access_token_expires"`
	RefreshTokenExpires time.Duration `json:"refresh_token_expires"`
}

func VerifyPassword(userPassword string, dbPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(userPassword))
}

func GetTokens(userId int, conf configuration.Auth) (*Tokens, error) {
	accessToken, err := CreateToken(userId, conf.AccessKey, conf.AccessExpire)
	if err != nil {
		return nil, err
	}

	refreshToken, err := CreateToken(userId, conf.AccessKey, conf.RefreshExpire)
	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:         accessToken,
		RefreshToken:        refreshToken,
		AccessTokenExpires:  conf.AccessExpire,
		RefreshTokenExpires: conf.RefreshExpire,
	}, nil
}

func CreateToken(userId int, key string, minute time.Duration) (string, error) {
	claims := authClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * minute).Unix(),
		},
	}
	tokensClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokensClaims.SignedString([]byte(key))
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
