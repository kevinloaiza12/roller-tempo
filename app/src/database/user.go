package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kevinloaiza12/roller-tempo/app/resources"
	_ "github.com/lib/pq"
)

// Creation

func CreateNewUser(ctx context.Context, db *sql.DB, data *resources.User) (bool, error) {

	id := data.GetUserID()
	monedas := data.GetUserCoins()
	turno := data.GetUserTurn()

	_, err := db.ExecContext(
		ctx,
		"INSERT INTO usuarios (id, monedas, turno) VALUES ($1,$2,$3)",
		id,
		monedas,
		turno,
	)

	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Getters

func GetUserByID(ctx context.Context, db *sql.DB, userID int) (*resources.User, error) {

	query := fmt.Sprintf(
		"SELECT %s,%s,%s FROM usuarios WHERE id = $1",
		"id",
		"monedas",
		"turno",
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
	} else {
		return resources.NewUser(
			id,
			coins,
			turn,
		), nil
	}
}

// Update

func UsersUpdateQuery(ctx context.Context, db *sql.DB, user *resources.User) (bool, error) {
	var query string
	query = fmt.Sprintf(
		"UPDATE usuarios SET monedas = %d, turno = %d "+
			"WHERE id = $1",
		user.GetUserCoins(),
		user.GetUserTurn(),
	)

	if _, err := db.ExecContext(ctx, query, user.GetUserID()); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
