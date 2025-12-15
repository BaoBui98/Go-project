package app

import (
	handler "user-management-user/internal/handller"
	"user-management-user/internal/repository"
	route "user-management-user/internal/routes"
	"user-management-user/internal/service"
)

type UserModule struct {
	route route.Router
}

func NewUserModule() *UserModule {
	userRepo := repository.NewUserRepository()

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	userRouter := route.NewUserRoute(userHandler)

	return &UserModule{
		route: userRouter,
	}
}

func (m *UserModule) GetRoute() route.Router {
	return m.route
}
