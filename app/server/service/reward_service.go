package service

import (
	"roller-tempo/model"
	"roller-tempo/repository"
)

type RewardService struct {
	rewardRepo *repository.RewardRepository
}

func NewRewardService(rewardRepo *repository.RewardRepository) *RewardService {
	return &RewardService{rewardRepo: rewardRepo}
}

func (rs *RewardService) CreateReward(reward *model.Reward) error {
	return rs.rewardRepo.CreateReward(reward)
}

func (rs *RewardService) UpdateReward(reward *model.Reward) error {
	return rs.rewardRepo.UpdateReward(reward)
}

func (rs *RewardService) DeleteReward(reward *model.Reward) error {
	return rs.rewardRepo.DeleteReward(reward)
}

func (rs *RewardService) GetRewardByID(id int) (*model.Reward, error) {
	return rs.rewardRepo.GetRewardById(id)
}
