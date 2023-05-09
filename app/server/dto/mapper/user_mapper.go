package dto

import (
	"roller-tempo/dto"
	"roller-tempo/model"
)

func ToUserDTO(user *model.User) *dto.UserDTO {
	userDTO := &dto.UserDTO{
		Identification: user.Identification,
		Coins:          user.Coins,
		Turn:           user.Turn,
		Attraction:     user.Attraction,
	}

	return userDTO
}
