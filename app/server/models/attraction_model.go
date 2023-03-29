package models

type Attraction struct {
	name        string
	description string
	duration    int
	capacity    int
  currentTurn int
	nextTurn    int
}

func (obj *Attraction) AttractionToJSON() map[string]interface{} {
	return map[string]interface{}{
		"name":        obj.name,
		"description": obj.description,
		"duration":    obj.duration,
		"capacity":    obj.capacity,
    "currentTurn": obj.currentTurn,
		"nextTurn":    obj.nextTurn,
	}
}

func NewAttraction(name string, description string, duration int, capacity int,currentTurn int, nextTurn int) *Attraction {
	return &Attraction{
		name,
		description,
		duration,
		capacity,
    currentTurn,
		nextTurn,
	}
}

// Setters

func (obj *Attraction) SetAttractionName(name string) {
	obj.name = name
}

func (obj *Attraction) SetAttractionDescription(description string) {
	obj.description = description
}

func (obj *Attraction) SetAttractionDuration(duration int) {
	obj.duration = duration
}

func (obj *Attraction) SetAttractionCapacity(capacity int) {
	obj.capacity = capacity
}

func (obj *Attraction) SetAttractionCurrentTurn(turn int) {
	obj.currentTurn = turn
}

func (obj *Attraction) SetAttractionNextTurn(turn int) {
	obj.nextTurn = turn
}

// Getters

func (obj *Attraction) GetAttractionName() string {
	return obj.name
}

func (obj *Attraction) GetAttractionDescription() string {
	return obj.description
}

func (obj *Attraction) GetAttractionDuration() int {
	return obj.duration
}

func (obj *Attraction) GetAttractionCapacity() int {
	return obj.capacity
}

func (obj *Attraction) GetAttractionCurrentTurn() int {
	return obj.currentTurn
}

func (obj *Attraction) GetAttractionNextTurn() int {
	return obj.nextTurn
}
