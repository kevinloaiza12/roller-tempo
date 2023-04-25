package models

type User struct {
	id    int
	coins int
	turn  int
    attraction string
}

func NewUser(id int, coins int, turn int, attraction string) *User {
	return &User{
		id,
		coins,
		turn,
        attraction,
	}
}

func (obj *User) UserToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":    obj.id,
		"coins": obj.coins,
		"turn":  obj.turn,
		"attraction":  obj.attraction,
	}
}

// Setters

func (obj *User) SetUserID(id int) {
	obj.id = id
}

func (obj *User) SetUserCoins(coins int) {
	obj.coins = coins
}

func (obj *User) SetUserTurn(turn int) {
	obj.turn = turn
}

func (obj *User) SetUserAttraction(attraction string) {
	obj.attraction = attraction
}

// Getters

func (obj *User) GetUserID() int {
	return obj.id
}

func (obj *User) GetUserCoins() int {
	return obj.coins
}

func (obj *User) GetUserTurn() int {
	return obj.turn
}

func (obj *User) GetUserAttraction() string {
	return obj.attraction
}
