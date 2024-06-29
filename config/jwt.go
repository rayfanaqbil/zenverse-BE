// package/config/jwt.go
package config

import (
	"fmt"
    "time"
    "github.com/golang-jwt/jwt/v4"
)

var SecretKey = []byte("ZeNvErSERynHrSZ")

func GenerateJWT(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString(SecretKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return SecretKey, nil
    })
}
