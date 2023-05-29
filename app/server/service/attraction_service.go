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

func (as *AttractionService) CreateAttraction(attraction *model.Attraction) error {
	return as.attractionRepository.CreateAttraction(attraction)
}

func (as *AttractionService) UpdateAttraction(attraction *model.Attraction) error {
	return as.attractionRepository.UpdateAttraction(attraction)
}

func (as *AttractionService) DeleteAttraction(attraction *model.Attraction) error {
	return as.attractionRepository.DeleteAttraction(attraction)
}

func (as *AttractionService) GetAllAttractions() ([]*model.Attraction, error) {
	return as.attractionRepository.GetAllAttractions()
}

func (as *AttractionService) GetAttractionByID(id int) (*model.Attraction, error) {
	return as.attractionRepository.GetAttractionByID(id)
}

func (as *AttractionService) GetNextAvailableTurn(attractionID int) (int, error) {
	attraction, err := as.attractionRepository.GetAttractionByID(attractionID)
	if err != nil {
		return 0, err
	}

	nextTurn := attraction.NextTurn

	return nextTurn, nil
}

func (as *AttractionService) GetNextRoundTurns(attractionID int) ([]int, error) {
	attraction, err := as.attractionRepository.GetAttractionByID(attractionID)
	if err != nil {
		return nil, err
	}

	nextRoundFiastTurn := (attraction.CurrentRoundTurn/attraction.Capacity+1)*attraction.Capacity + 1

	nextRoundTurns := []int{}

	for i := nextRoundFiastTurn; i < nextRoundFiastTurn+attraction.Capacity; i++ {
		nextRoundTurns = append(nextRoundTurns, i)
	}

	return nextRoundTurns, nil
}

func (as *AttractionService) UpdateAttractionTurns(attractionID int) error {
	attraction, err := as.attractionRepository.GetAttractionByID(attractionID)
	if err != nil {
		return err
	}

	attraction.CurrentRoundTurn = attraction.CurrentRoundTurn + attraction.Capacity

	return as.attractionRepository.UpdateAttraction(attraction)
}
