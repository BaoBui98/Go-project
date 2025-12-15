package route

import (
	handler "user-management-user/internal/handller"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	handler *handler.UserHandler
}

func NewUserRoute(handler *handler.UserHandler) *UserRoute {
	return &UserRoute{
		handler: handler,
	}
}
func (ur *UserRoute) Register(r *gin.RouterGroup) {
	user := r.Group("/users")
	{
		user.GET("/", ur.handler.GetAllUsers)
		user.POST("/", ur.handler.CreateUser)
		user.GET("/:uuid", ur.handler.GetUserByUUID)
		user.PUT("/:uuid", ur.handler.UpdateUser)
		user.DELETE("/:uuid", ur.handler.DeleteUser)
	}
}
