package repository

import (
	"roller-tempo/model"

	"gorm.io/gorm"
)

type RewardRepository struct {
	db *gorm.DB
}

func NewRewardRepository(db *gorm.DB) *RewardRepository {
	return &RewardRepository{db: db}
}

func (ur *RewardRepository) GetAllRewards() ([]*model.Reward, error) {
	var rewards []*model.Reward
	err := ur.db.Find(&rewards).Error
	if err != nil {
		return nil, err
	}

	return rewards, nil
}

func (ur *RewardRepository) CreateReward(reward *model.Reward) error {
	return ur.db.Create(reward).Error
}

func (ur *RewardRepository) UpdateReward(reward *model.Reward) error {
	return ur.db.Save(reward).Error
}

func (ur *RewardRepository) DeleteReward(reward *model.Reward) error {
	return ur.db.Delete(reward).Error
}

func (ur *RewardRepository) GetRewardByID(id int) (*model.Reward, error) {
	var reward model.Reward
	err := ur.db.First(&reward, id).Error
	if err != nil {
		return nil, err
	}
	return &reward, nil
}
