package service

import (
	"user-management-user/internal/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}

}

func (us *userService) GetAllUsers() {
	us.repo.FindAll()
}

func (us *userService) CreateUser() {
	us.repo.Create()
}

func (us *userService) GetUserByUUID() {
	us.repo.GetByID()
}

func (us *userService) UpdateUser() {
	us.repo.Update()
}

func (us *userService) DeleteUser() {
	us.repo.Delete()
}
