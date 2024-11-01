package config

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rayfanaqbil/zenverse-BE/v2/model"
)

func GenerateJWT(admin model.Admin) (string, error) {
	claims := jwt.MapClaims{}
	claims["admin_id"] = admin.ID.Hex()
	claims["email"] = admin.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token valid for 24 hours


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secretKey))
}
