package models

type Reward struct {
	name        string
	description string
	price       int
}

func (obj *Reward) RewardToJSON() map[string]interface{} {
	return map[string]interface{}{
		"name":        obj.name,
		"description": obj.description,
		"price":       obj.price,
	}
}

func NewReward(name string, description string, price int) *Reward {
	return &Reward{
		name,
		description,
		price,
	}
}

// Setters

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

func (obj *Reward) GetRewardName() string {
	return obj.name
}

func (obj *Reward) GetRewardDescription() string {
	return obj.description
}

func (obj *Reward) GetRewardPrice() int {
	return obj.price
}
