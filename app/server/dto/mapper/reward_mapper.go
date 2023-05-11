package dto

import (
	"roller-tempo/dto"
	"roller-tempo/model"
)

func ToRewardDTO(reward *model.Reward) *dto.RewardDTO {
	rewardDTO := &dto.RewardDTO{
		Name:        reward.Name,
		Description: reward.Description,
		Price:       reward.Price,
		ImagePath:   reward.ImagePath,
	}

	return rewardDTO
}
