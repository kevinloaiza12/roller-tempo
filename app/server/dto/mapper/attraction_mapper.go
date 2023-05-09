package dto

import (
	"roller-tempo/dto"
	"roller-tempo/model"
)

func ToAttractionDTO(attraction *model.Attraction) *dto.AttractionDTO {
	attractionDTO := &dto.AttractionDTO{
		Name:             attraction.Name,
		Description:      attraction.Description,
		Duration:         attraction.Duration,
		Capacity:         attraction.Capacity,
		CurrentRoundTurn: attraction.CurrentRoundTurn,
		NextTurn:         attraction.NextTurn,
		PosX:             attraction.PosX,
		PosY:             attraction.PosY,
	}

	return attractionDTO
}
