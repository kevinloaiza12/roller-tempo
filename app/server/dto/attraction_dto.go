package dto

type AttractionDTO struct {
	Name        string  `gorm:"column:name"`
	Description string  `gorm:"column:description"`
	Duration    int     `gorm:"column:duration"`
	Capacity    int     `gorm:"column:capacity"`
	CurrentTurn int     `gorm:"column:currentTurn"`
	NextTurn    int     `gorm:"column:nextTurn"`
	X           float64 `gorm:"column:x"`
	Y           float64 `gorm:"column:y"`
}
