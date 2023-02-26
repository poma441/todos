package jwtparser

import (
	"encoding/base64"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func RsaSignedTokenParse(token string, publicKey string) (*jwt.Token, error) {
	decodePublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodePublicKey)
	if err != nil {
		return nil, fmt.Errorf("validate: ошибка: %w", err)
	}

	tokenParsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("непредвиденный метод: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	return tokenParsed, nil
}
