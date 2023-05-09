package controller

type CreateUserRequest struct {
	Identification int `json:"identification"`
	Coins          int `json:"coins"`
	Turn           int `json:"turn"`
	Attraction     int `json:"attraction"`
}

type UpdateUserTurnRequest struct {
	UserID       int `json:"userID"`
	AttractionID int `json:"attractionID"`
}
