package handler

import (
	"log"
	"user-management-user/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
	log.Println("GetAllUsers called")
	uh.service.GetAllUsers()
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	log.Println("CreateUser called")
	uh.service.CreateUser()
}

func (uh *UserHandler) GetUserByUUID(ctx *gin.Context) {
	log.Println("GetUserByUUID called")
	uh.service.GetUserByUUID()
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	log.Println("UpdateUser called")
	uh.service.UpdateUser()
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	log.Println("DeleteUser called")
	uh.service.DeleteUser()
}
