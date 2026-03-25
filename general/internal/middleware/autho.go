package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"general/internal/auth"

	"github.com/golang-jwt/jwt/v5"
)

type tenantKey struct{}

func AuthMiddleware(secret []byte) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("AuthMiddleware: request %s %s", r.Method, r.URL.Path)

			// Allow health checks
			if r.URL.Path == "/health" {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")
			log.Printf("AuthMiddleware: Authorization header = '%s'", authHeader)

			if authHeader == "" {
				log.Printf("AuthMiddleware: missing auth header")
				http.Error(w, `{"error":"missing authorization header"}`, http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			log.Printf("AuthMiddleware: parts count = %d", len(parts))
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				log.Printf("AuthMiddleware: invalid header format")
				http.Error(w, `{"error":"invalid authorization header format"}`, http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]
			log.Printf("AuthMiddleware: token (first 20 chars) = %s...", tokenString[:20])

			var claims auth.Claims
			token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return secret, nil
			})

			if err != nil {
				log.Printf("AuthMiddleware: JWT parse error: %v", err)
				http.Error(w, `{"error":"invalid token"}`, http.StatusUnauthorized)
				return
			}

			if !token.Valid {
				log.Printf("AuthMiddleware: token not valid")
				http.Error(w, `{"error":"invalid token"}`, http.StatusUnauthorized)
				return
			}

			if claims.TenantID == "" {
				log.Printf("AuthMiddleware: missing tenant_id claim")
				http.Error(w, `{"error":"missing tenant_id"}`, http.StatusUnauthorized)
				return
			}

			log.Printf("AuthMiddleware: tenantID = %s", claims.TenantID)
			ctx := context.WithValue(r.Context(), tenantKey{}, claims.TenantID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetTenantID(ctx context.Context) string {
	if tenantID, ok := ctx.Value(tenantKey{}).(string); ok {
		return tenantID
	}
	return ""
}
