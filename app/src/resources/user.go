package resources

type User struct {
	id    int
	coins int
	turn  int
}

func NewUser(id int, coins int, turn int) *User {
	return &User{
		id,
		coins,
		turn,
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
