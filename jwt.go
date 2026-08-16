package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenConfig struct {
	TokenSecret []byte
	TokenTTL    time.Duration
	Issuer      string
}

type UserClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func NewConfig() (*TokenConfig, error) {

	accessSecret := os.Getenv("JWT_SECRET")

	if len(accessSecret) < 32 {
		return nil, errors.New("JWT secret must be at least 32 characters")
	}

	return &TokenConfig{
		TokenSecret: []byte(accessSecret),
		TokenTTL:    20 * time.Minute,
		Issuer:      "auth-api",
	}, nil
}

func CreateToken(cfg *TokenConfig, userID string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": cfg.Issuer,
		"sub": userID,
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(time.Now().Add(cfg.TokenTTL)),
	})

	return token.SignedString(cfg.TokenSecret)
}

func ValidateToken(tokenString string, secret []byte) (*UserClaims, error) {

	claims := &UserClaims{}

	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("signing method in jwt unepexpected", "alg", t.Header["alg"])
			return nil, fmt.Errorf("signing method unepexpected")
		}

		return []byte(secret), nil
	})

	if err != nil || !parsedToken.Valid {
		// INVALID TOKEN OR EXPIRED
		return nil, err
	}

	return claims, nil
}
