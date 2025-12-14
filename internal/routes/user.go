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

}
