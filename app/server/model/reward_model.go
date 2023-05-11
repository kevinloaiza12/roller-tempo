package model

import "gorm.io/gorm"

type Reward struct {
	gorm.Model

	Name        string `gorm:"column:name;unique"`
	Description string `gorm:"column:description"`
	Price       int    `gorm:"column:price"`
	ImagePath   string `gorm:"column:image_path"`
}
