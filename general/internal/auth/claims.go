package auth

import "github.com/golang-jwt/jwt/v5"

// Claims is the standard JWT structure used across Asguard
type Claims struct {
	TenantID string `json:"tenant_id"`
	jwt.RegisteredClaims
}