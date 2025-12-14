package repository

import "user-management-user/internal/model"

type InMemoryUserRepository struct {
	users []model.Users
}

func NewUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make([]model.Users, 0),
	}
}
