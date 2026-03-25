package middleware

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter interface {
	Allow(ctx context.Context, key string) bool
}

type rateClient struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type LocalRateLimiter struct {
	mu       sync.RWMutex
	limiters map[string]*rateClient
	rate     rate.Limit
	burst    int
}

func NewLocalRateLimiter(rps float64, burst int) *LocalRateLimiter {
	limiter := &LocalRateLimiter{
		rate:     rate.Limit(rps),
		burst:    burst,
		limiters: make(map[string]*rateClient),
	}
	go limiter.cleanup() // Start cleanup goroutine
	return limiter
}

func (l *LocalRateLimiter) Allow(ctx context.Context, key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	client, exists := l.limiters[key]
	if !exists {
		client = &rateClient{
			limiter:  rate.NewLimiter(l.rate, l.burst),
			lastSeen: time.Now(),
		}
		l.limiters[key] = client
	}

	client.lastSeen = time.Now()
	return client.limiter.Allow()
}

func (l *LocalRateLimiter) cleanup() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		l.mu.Lock()
		for key, client := range l.limiters {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(l.limiters, key)
			}
		}
		l.mu.Unlock()
	}
}

func RateLimitMiddleware(limiter RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("RateLimit: tenantID=%q from context", GetTenantID(r.Context()))
			tenantID := GetTenantID(r.Context())
			if tenantID == "" {
				next.ServeHTTP(w, r)
				return
			}

			if !limiter.Allow(r.Context(), tenantID) {
				http.Error(w, `{"error":"rate limit exceeded"}`, http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
