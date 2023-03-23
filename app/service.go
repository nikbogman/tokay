package app

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignToken(subject, issuer, secret string, lifeTime time.Duration) (string, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Duration(lifeTime)))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   subject,
		ExpiresAt: expiresAt,
	})
	signed, err := token.SignedString([]byte(secret))
	return signed, err
}

func DecodeToken(token, secret string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	decoded, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if !decoded.Valid {
		err := errors.New("token is not valid")
		return nil, err
	}
	return claims, err
}
