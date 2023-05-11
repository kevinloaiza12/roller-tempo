package model

import "gorm.io/gorm"

type Attraction struct {
	gorm.Model

	Name             string  `gorm:"column:name;unique"`
	Description      string  `gorm:"column:description"`
	Duration         int     `gorm:"column:duration"`
	Capacity         int     `gorm:"column:capacity"`
	CurrentRoundTurn int     `gorm:"column:current_round_turn"`
	NextTurn         int     `gorm:"column:next_turn"`
	PosX             float64 `gorm:"column:x"`
	PosY             float64 `gorm:"column:y"`
	ImagePath        string  `gorm:"column:image_path"`
}
