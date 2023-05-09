package model

import "gorm.io/gorm"

type Reward struct {
	gorm.Model

	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Price       int    `gorm:"column:price"`
}

func (r *Reward) GetName() string {
	return r.Name
}

func (r *Reward) GetDescription() string {
	return r.Description
}

func (r *Reward) GetPrice() int {
	return r.Price
}

func (r *Reward) SetName(name string) {
	r.Name = name
}

func (r *Reward) SetDescription(description string) {
	r.Description = description
}

func (r *Reward) setPrice(price int) {
	r.Price = price
}
