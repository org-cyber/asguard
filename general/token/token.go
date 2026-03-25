package main

import (
	"fmt"
	"os"
	"time"

	"general/internal/auth"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func main() {
	// token.go
	_ = godotenv.Load()
	secret := []byte(os.Getenv("JWT_SECRET")) // Match your JWT_SECRET env var

	claims := auth.Claims{
		TenantID: "tenant-acme-corp",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	fmt.Println(tokenString)
}
