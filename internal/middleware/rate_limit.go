package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*Client)
)

func getClientIP(ctx *gin.Context) string {
	return ctx.ClientIP()
}

func getRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	client, exists := clients[ip]
	if !exists {
		limiter := rate.NewLimiter(5, 10) // 5 requests per second with a burst of 10
		clients[ip] = &Client{limiter: limiter, lastSeen: time.Now()}
		return limiter
	}
	client.lastSeen = time.Now()

	return client.limiter

}

func RateLimitMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIP(ctx)
		limiter := getRateLimiter(ip)
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(429, gin.H{"error": "Too Many Requests"})
			return
		}
		ctx.Next()
	}
}

func cleanupClients() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}
