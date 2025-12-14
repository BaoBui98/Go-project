package service

import "user-management-user/internal/repository"

type UserService struct {
	repo *repository.InMemoryUserRepository
}

func NewUserService(repo *repository.InMemoryUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}

}
