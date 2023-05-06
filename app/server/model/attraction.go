package models

import "gorm.io/gorm"

type Attraction struct {
	gorm.Model

	name        string
	description string
	duration    int
	capacity    int
	currentTurn int
	nextTurn    int
	x           float64
	y           float64
}
