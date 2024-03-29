package repository

import (
	"roller-tempo/model"

	"gorm.io/gorm"
)

type AttractionRepository struct {
	db *gorm.DB
}

func NewAttractionRepository(db *gorm.DB) *AttractionRepository {
	return &AttractionRepository{db: db}
}

func (ur *AttractionRepository) GetAllAttractions() ([]*model.Attraction, error) {
	var attractions []*model.Attraction
	err := ur.db.Find(&attractions).Error
	if err != nil {
		return nil, err
	}

	return attractions, nil
}

func (ur *AttractionRepository) CreateAttraction(attraction *model.Attraction) error {
	return ur.db.Create(attraction).Error
}

func (ur *AttractionRepository) UpdateAttraction(attraction *model.Attraction) error {
	return ur.db.Save(attraction).Error
}

func (ur *AttractionRepository) DeleteAttraction(attraction *model.Attraction) error {
	return ur.db.Delete(attraction).Error
}

func (ur *AttractionRepository) GetAttractionByID(id int) (*model.Attraction, error) {
	var attraction model.Attraction
	err := ur.db.First(&attraction, id).Error
	if err != nil {
		return nil, err
	}
	return &attraction, nil
}
