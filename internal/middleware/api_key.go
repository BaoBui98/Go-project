package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ApiKeyMiddleware() gin.HandlerFunc {
	expectedAPIKey := os.Getenv("API_KEY")
	if expectedAPIKey == "" {
		expectedAPIKey = "secret-api-key"
	}
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-API-KEY")
		if apiKey == "" {
			ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": "API key required"})
			return
		}
		if apiKey != expectedAPIKey {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		ctx.Next()
	}
}
