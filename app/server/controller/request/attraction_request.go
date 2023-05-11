package controller

type CreateAttractionRequest struct {
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Duration         int     `json:"duration"`
	Capacity         int     `json:"capacity"`
	CurrentRoundTurn int     `json:"currentRoundTurn"`
	NextTurn         int     `json:"nextTurn"`
	PosX             float64 `json:"x"`
	PosY             float64 `json:"y"`
}
