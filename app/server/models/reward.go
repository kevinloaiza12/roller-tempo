package models

type Reward struct {
	id          int64
	name        string
	description string
	price       int
}

func (obj *Reward) RewardToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":          obj.id,
		"name":        obj.name,
		"description": obj.description,
		"price":       obj.price,
	}
}

func NewReward(id int64, name string, description string, price int) *Reward {
	return &Reward{
		id,
		name,
		description,
		price,
	}
}

func (u *User) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":    u.id,
		"coins": u.coins,
		"turn":  u.turn,
	}
}

// Setters

func (obj *Reward) SetRewardID(id int64) {
	obj.id = id
}

func (obj *Reward) SetRewardName(name string) {
	obj.name = name
}

func (obj *Reward) SetRewardDescription(description string) {
	obj.description = description
}

func (obj *Reward) SetRewardPrice(price int) {
	obj.price = price
}

// Getters

func (obj *Reward) GetRewardID() int64 {
	return obj.id
}

func (obj *Reward) GetRewardName() string {
	return obj.name
}

func (obj *Reward) GetRewardDescription() string {
	return obj.description
}

func (obj *Reward) GetRewardPrice() int {
	return obj.price
}
