package dto

type RewardDTO struct {
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Price       int    `gorm:"column:price"`
}
