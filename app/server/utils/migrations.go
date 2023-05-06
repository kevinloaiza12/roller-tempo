package utils

import (
	model "roller-tempo/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(model.Attraction{})
	db.AutoMigrate(model.Reward{})
	db.AutoMigrate(model.User{})
}
