package dto

type AttractionDTO struct {
	Name             string
	Description      string
	Duration         int
	Capacity         int
	CurrentRoundTurn int
	NextTurn         int
	PosX             float64
	PosY             float64
}
