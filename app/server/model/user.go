package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	id         int
	coins      int
	turn       int
	attraction string
}
