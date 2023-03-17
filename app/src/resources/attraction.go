package resources

type Attraction struct {
	id          int64
	name        string
	description string
	duration    int
	capacity    int
	nextTurn    int
}

func NewAttraction(id int64, name string, description string, duration int, capacity int, nextTurn int) *Attraction {
	return &Attraction{
		id,
		name,
		description,
		duration,
		capacity,
		nextTurn,
	}
}

// Setters

func (obj *Attraction) SetAttractionID(id int64) {
	obj.id = id
}

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

func (obj *Attraction) SetAttractionNextTurn(turn int) {
	obj.nextTurn = turn
}

// Getters

func (obj *Attraction) GetAttractionID() *int64 {
	return &obj.id
}

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

func (obj *Attraction) GetAttractionNextTurn() int {
	return obj.nextTurn
}
