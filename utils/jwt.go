package utils

import (
	"github.com/dgrijalva/jwt-go"
	"wserver/conf"
)

var secret []byte

func init() {
	secret = []byte(conf.AppConf.Server.WsAuthKey)
}

func MakeToken(claims jwt.StandardClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	return tokenString, err
}

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
