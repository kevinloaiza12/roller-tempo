package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Identification int
	Coins          int
	Turn           int
	Attraction     int
}

func (u *User) GetIdentification() int {
	return u.Identification
}

func (u *User) SetIdentification(identification int) {
	u.Identification = identification
}

func (u *User) GetCoins() int {
	return u.Coins
}

func (u *User) SetCoins(value int) {
	u.Coins = value
}

func (u *User) GetTurn() int {
	return u.Turn
}

func (u *User) SetTurn(value int) {
	u.Turn = value
}

func (u *User) GetAttraction() int {
	return u.Attraction
}

func (u *User) SetAttraction(attractionID int) {
	u.Attraction = attractionID
}
