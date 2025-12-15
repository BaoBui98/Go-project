package app

import (
	"log"
	"user-management-user/internal/config"
	route "user-management-user/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Module interface {
	GetRoute() route.Router
}

type Application struct {
	config  *config.Config
	router  *gin.Engine
	modules []Module
}

func NewApplication(cfg *config.Config) *Application {
	r := gin.Default()
	modules := []Module{
		NewUserModule(),
	}
	route.RegisterRoute(r, getModuleRoutes(modules)...)
	return &Application{
		config:  cfg,
		router:  r,
		modules: modules,
	}
}

func (a *Application) Run() error {
	return a.router.Run(a.config.ServerAddress)
}

func getModuleRoutes(modules []Module) []route.Router {
	routes := make([]route.Router, 0, len(modules))
	for _, module := range modules {
		routes = append(routes, module.GetRoute())
	}
	return routes
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}
