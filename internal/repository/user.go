package repository

import (
	"log"
	"user-management-user/internal/model"
)

type InMemoryUserRepository struct {
	users []model.Users
}

func NewUserRepository() UserRepository {
	return &InMemoryUserRepository{
		users: make([]model.Users, 0),
	}
}

func (ur *InMemoryUserRepository) FindAll() {
	log.Println("GetAllUsers repo called")
}

func (ur *InMemoryUserRepository) Create() {
	log.Println("CreateUser repo called")
}

func (ur *InMemoryUserRepository) GetByID() {
	log.Println("GetUserByUUID repo called")
}

func (ur *InMemoryUserRepository) Update() {
	log.Println("UpdateUser repo called")
}

func (ur *InMemoryUserRepository) Delete() {
	log.Println("DeleteUser repo called")
}
