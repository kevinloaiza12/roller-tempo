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

	_, err := db.ExecContext(
		ctx,
		"INSERT INTO users (UserID, UserCoins, UserTurn) VALUES ($1,$2,$3)",
		id,
		coins,
		turn,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

// Getters

func GetUserByID(ctx context.Context, db *sql.DB, userID int) (*models.User, error) {

	query := fmt.Sprintf(
		"SELECT %s,%s,%s FROM users WHERE UserID = $1",
		"UserID",
		"UserCoins",
		"UserTurn",
	)

	var id int
	var coins int
	var turn int

	err := db.QueryRowContext(ctx, query, userID).Scan(
		&id,
		&coins,
		&turn,
	)

	if err != nil {
		return nil, err
	}

	return models.NewUser(
		id,
		coins,
		turn,
	), nil
}

// Update

func UsersUpdateQuery(ctx context.Context, db *sql.DB, user *models.User) (bool, error) {

	query := fmt.Sprintf(
		"UPDATE users SET UserCoins = %d, UserTurn = %d "+
			"WHERE UserID = $1",
		user.GetUserCoins(),
		user.GetUserTurn(),
	)

	if _, err := db.ExecContext(ctx, query, user.GetUserID()); err != nil {
		return false, err
	}

	return true, nil
}
