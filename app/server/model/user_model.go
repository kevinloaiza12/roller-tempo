package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Identification int `gorm:"column:identification;unique"`
	Coins          int `gorm:"column:coins"`
	Turn           int `gorm:"column:turn"`
	Attraction     int `gorm:"column:attraction"`
}
