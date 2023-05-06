package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	identification int
	coins          int
	turn           int
	attraction     string
}
