package repository

import (
	"roller-tempo/dto"
	"roller-tempo/model"

	"gorm.io/gorm"
)

type RewardRepository struct {
	db *gorm.DB
}

func NewRewardRepository(db *gorm.DB) *RewardRepository {
	return &RewardRepository{db: db}
}

func (ur *RewardRepository) GetAllRewards() ([]*dto.RewardDTO, error) {
	var rewards []*model.Reward
	err := ur.db.Find(&rewards).Error
	if err != nil {
		return nil, err
	}

	var rewardDTOs []*dto.RewardDTO
	for _, reward := range rewards {
		rewardDTO := &dto.RewardDTO{
			Name:        reward.Name,
			Description: reward.Description,
			Price:       reward.Price,
		}
		rewardDTOs = append(rewardDTOs, rewardDTO)
	}

	return rewardDTOs, nil
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
