package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kevinloaiza12/roller-tempo/app/models"
	_ "github.com/lib/pq"
)

// Creation

func CreateNewUser(ctx context.Context, db *sql.DB, data *models.User) (bool, error) {

	id := data.GetUserID()
	coins := data.GetUserCoins()
	turn := data.GetUserTurn()
    attraction := data.GetUserAttraction()

	_, err := db.ExecContext(
		ctx,
		"INSERT INTO users (UserID, UserCoins, UserTurn, UserCurrentAttraction) VALUES ($1,$2,$3,$4)",
		id,
		coins,
		turn,
        attraction,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

// Getters

func GetUserByID(ctx context.Context, db *sql.DB, userID int) (*models.User, error) {

	query := fmt.Sprintf(
		"SELECT %s,%s,%s,%s FROM users WHERE UserID = $1",
		"UserID",
		"UserCoins",
		"UserTurn",
        "UserCurrentAttraction",
	)

	var id int
	var coins int
	var turn int
    var attraction string

	err := db.QueryRowContext(ctx, query, userID).Scan(
		&id,
		&coins,
		&turn,
        &attraction,
	)

	if err != nil {
		return nil, err
	}

	return models.NewUser(
		id,
		coins,
		turn,
        attraction,
	), nil
}

// Update

func UsersUpdateQuery(ctx context.Context, db *sql.DB, user *models.User) (bool, error) {

	query := fmt.Sprintf(
		"UPDATE users SET UserCoins = %d, UserTurn = %d, UserCurrentAttraction = '%s' "+
			"WHERE UserID = $1",
		user.GetUserCoins(),
		user.GetUserTurn(),
        user.GetUserAttraction(),
	)

	if _, err := db.ExecContext(ctx, query, user.GetUserID()); err != nil {
		return false, err
	}

	return true, nil
}

