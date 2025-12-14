package main

import (
	"user-management-user/internal/config"
	handler "user-management-user/internal/handller"
	"user-management-user/internal/repository"
	route "user-management-user/internal/routes"
	"user-management-user/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	userRepo := repository.NewUserRepository()

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	userRouter := route.NewUserRoute(userHandler)

	r := gin.Default()
	route.RegisterRoute(r, userRouter)
	if err := r.Run(cfg.ServerAddress); err != nil {
		panic(err)
	}

}
