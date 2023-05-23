package controller

type CreateUserRequest struct {
	Identification int `json:"identification"`
	Coins          int `json:"coins"`
}

type UpdateUserTurnRequest struct {
	UserID       int `json:"userID"`
	AttractionID int `json:"attractionID"`
}

type BuyRewardRequest struct {
	UserID   int `json:"userID"`
	RewardID int `json:"rewardID"`
}
