package main

import (
	"user-management-user/internal/app"
	"user-management-user/internal/config"
)

func main() {
	cfg := config.NewConfig()

	application := app.NewApplication(cfg)

	if err := application.Run(); err != nil {
		panic(err)
	}

}
