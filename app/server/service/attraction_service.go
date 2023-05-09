package service

import (
	"roller-tempo/model"
	"roller-tempo/repository"
)

type AttractionService struct {
	attractionRepo *repository.AttractionRepository
}

func NewAttractionService(attractionRepo *repository.AttractionRepository) *AttractionService {
	return &AttractionService{attractionRepo: attractionRepo}
}

func (rs *AttractionService) CreateAttraction(attraction *model.Attraction) error {
	return rs.attractionRepo.CreateAttraction(attraction)
}

func (rs *AttractionService) UpdateAttraction(attraction *model.Attraction) error {
	return rs.attractionRepo.UpdateAttraction(attraction)
}

func (rs *AttractionService) DeleteAttraction(attraction *model.Attraction) error {
	return rs.attractionRepo.DeleteAttraction(attraction)
}

func (rs *AttractionService) GetAttractionByID(id int) (*model.Attraction, error) {
	return rs.attractionRepo.GetAttractionByID(id)
}
