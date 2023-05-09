package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Identification int
	Coins          int
	Turn           int
	Attraction     int
}
