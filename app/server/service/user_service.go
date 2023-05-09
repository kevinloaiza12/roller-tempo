package service

import (
	"roller-tempo/model"
	"roller-tempo/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) CreateUser(user *model.User) error {
	return us.userRepo.CreateUser(user)
}

func (us *UserService) UpdateUser(user *model.User) error {
	return us.userRepo.UpdateUser(user)
}

func (us *UserService) DeleteUser(user *model.User) error {
	return us.userRepo.DeleteUser(user)
}

func (us *UserService) GetUserByID(id int) (*model.User, error) {
	return us.userRepo.GetUserByID(id)
}
