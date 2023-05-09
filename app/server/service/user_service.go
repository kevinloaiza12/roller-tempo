package service

import (
	"roller-tempo/model"
	"roller-tempo/repository"
)

type UserService struct {
	userRepository    *repository.UserRepository
	attractionService *AttractionService
}

func NewUserService(userRepo *repository.UserRepository, attractionService *AttractionService) *UserService {
	return &UserService{
		userRepository:    userRepo,
		attractionService: attractionService,
	}
}

func (us *UserService) CreateUser(user *model.User) error {
	return us.userRepository.CreateUser(user)
}

func (us *UserService) UpdateUser(user *model.User) error {
	return us.userRepository.UpdateUser(user)
}

func (us *UserService) DeleteUser(user *model.User) error {
	return us.userRepository.DeleteUser(user)
}

func (us *UserService) GetUserByID(id int) (*model.User, error) {
	return us.userRepository.GetUserByID(id)
}

func (us *UserService) UpdateUserTurnAndAttraction(id int, attractionID int) error {
	user, err := us.userRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	attraction, err := us.attractionService.GetAttractionByID(attractionID)
	if err != nil {
		return err
	}

	user.Turn = attraction.NextTurn
	user.Attraction = attractionID

	err = us.userRepository.UpdateUser(user)
	if err != nil {
		return err
	}

	attraction.NextTurn += 1
	err = us.attractionService.UpdateAttraction(attraction)
	if err != nil {
		return err
	}

	return nil
}
