package route

import (
	"user-management-user/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoute(r *gin.Engine, routes ...Router) {
	r.Use(middleware.AuthMiddleware())
	api := r.Group("/api/v1")
	for _, route := range routes {
		route.Register(api)
	}
}
