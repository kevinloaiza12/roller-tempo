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

func (us *UserService) GetAllUsers() ([]*model.User, error) {
	return us.userRepository.GetAllUsers()
}

func (us *UserService) GetUserByID(id int) (*model.User, error) {
	return us.userRepository.GetUserByID(id)
}

func (us *UserService) UpdateUserTurnAndAttraction(userID int, attractionID int) error {
	user, err := us.userRepository.GetUserByID(userID)
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

func (us *UserService) ValidateTurn(userID int, attractionID int) (bool, error) {
	user, err := us.userRepository.GetUserByID(userID)
	if err != nil {
		return false, err
	}

	if user.Attraction != attractionID {
		return false, nil
	}

	turns, err := us.attractionService.GetNextRoundTurns(attractionID)
	if err != nil {
		return false, err
	}

	containsValue := func(slice []int, value int) bool {
		for _, item := range slice {
			if item == value {
				return true
			}
		}
		return false
	}

	if containsValue(turns, user.Turn) {
		return true, nil
	}

	return false, nil
}

func (us *UserService) RemoveTurn(id int) error {
	user, err := us.userRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	user.Turn = 0

	return us.UpdateUser(user)
}

func (us *UserService) RewardUser(userID int, amount int) error {
	user, err := us.userRepository.GetUserByID(userID)
	if err != nil {
		return err
	}

	user.Coins += amount

	return us.UpdateUser(user)
}

func (us *UserService) PenalizeUser(userID int, amount int) error {
	user, err := us.userRepository.GetUserByID(userID)
	if err != nil {
		return err
	}

	user.Coins -= amount

	if user.Coins < 0 {
		user.Coins = 0
	}

	return us.UpdateUser(user)
}
