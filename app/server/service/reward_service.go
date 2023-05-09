package service

import (
	"roller-tempo/model"
	"roller-tempo/repository"
)

type RewardService struct {
	rewardRepository *repository.RewardRepository
}

func NewRewardService(rewardRepo *repository.RewardRepository) *RewardService {
	return &RewardService{rewardRepository: rewardRepo}
}

func (rs *RewardService) CreateReward(reward *model.Reward) error {
	return rs.rewardRepository.CreateReward(reward)
}

func (rs *RewardService) UpdateReward(reward *model.Reward) error {
	return rs.rewardRepository.UpdateReward(reward)
}

func (rs *RewardService) DeleteReward(reward *model.Reward) error {
	return rs.rewardRepository.DeleteReward(reward)
}

func (rs *RewardService) GetRewardByID(id int) (*model.Reward, error) {
	return rs.rewardRepository.GetRewardByID(id)
}
