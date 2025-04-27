package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "Backend"

func GenerateJwt(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
func VerifyToken(token string) (*int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errors.New("could not parse token")
	}
	isTokenValid := parsedToken.Valid
	if !isTokenValid {
		return nil, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	// exp, ok := claims["exp"].(int64)
	// fmt.Println(ok)
	fmt.Println(email, userId)
	return &userId, nil
}
