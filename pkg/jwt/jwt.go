package jwtToken

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

var Secret = "SECRET_KEY"

func GeneratorToken(claim *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	webToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "gagal signed", err
	}

	return webToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}


func DecodeToken(tokenString string)(jwt.MapClaims, error){
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims , isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}








