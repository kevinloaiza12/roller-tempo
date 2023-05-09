package service

import (
	"roller-tempo/model"
	"roller-tempo/repository"
)

type AttractionService struct {
	attractionRepository *repository.AttractionRepository
}

func NewAttractionService(attractionRepo *repository.AttractionRepository) *AttractionService {
	return &AttractionService{attractionRepository: attractionRepo}
}

func (rs *AttractionService) CreateAttraction(attraction *model.Attraction) error {
	return rs.attractionRepository.CreateAttraction(attraction)
}

func (rs *AttractionService) UpdateAttraction(attraction *model.Attraction) error {
	return rs.attractionRepository.UpdateAttraction(attraction)
}

func (rs *AttractionService) DeleteAttraction(attraction *model.Attraction) error {
	return rs.attractionRepository.DeleteAttraction(attraction)
}

func (rs *AttractionService) GetAttractionByID(id int) (*model.Attraction, error) {
	return rs.attractionRepository.GetAttractionByID(id)
}

func (as *AttractionService) GetNextAvailableTurn(attractionID int) (int, error) {
	attraction, err := as.attractionRepository.GetAttractionByID(attractionID)
	if err != nil {
		return 0, err
	}

	nextTurn := attraction.CurrentRoundTurn + 1

	return nextTurn, nil
}

func (as *AttractionService) GetNextRoundTurns(attractionID int) ([]int, error) {

	attraction, err := as.attractionRepository.GetAttractionByID(attractionID)
	if err != nil {
		return nil, err
	}

	nextRoundFirstTurn := (attraction.CurrentRoundTurn/attraction.Capacity + 1) * attraction.Capacity

	nextRoundTurns := []int{}

	for i := nextRoundFirstTurn; i < nextRoundFirstTurn+attraction.Capacity; i++ {
		nextRoundTurns = append(nextRoundTurns, i)
	}

	return nextRoundTurns, nil
}
