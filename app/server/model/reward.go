package models

import "gorm.io/gorm"

type Reward struct {
	gorm.Model

	name        string
	description string
	price       int
}
