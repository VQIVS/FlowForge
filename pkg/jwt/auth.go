package jwt

import (
	"errors"

	jwt2 "github.com/golang-jwt/jwt/v5"
)

const UserClaimKey = "User-Claim"

func CreateToken(secret []byte, claims *UserClaims) (string, error) {
	return jwt2.NewWithClaims(jwt2.SigningMethodHS256, claims).SignedString(secret)
}

func ParseToken(tokenString string, secret []byte) (*UserClaims, error) {
	token, err := jwt2.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt2.Token) (interface{}, error) {
		return secret, nil
	})
	if token == nil {
		return nil, errors.New("invalid token")
	}

	var claim *UserClaims
	if token.Claims != nil {
		cc, ok := token.Claims.(*UserClaims)
		if ok {
			claim = cc
		}
	}
	if err != nil {
		return claim, err
	}

	if !token.Valid {
		return claim, errors.New("invalid token")
	}
	return claim, nil
}
