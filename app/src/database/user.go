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

// Utils

func usersGetQuery(ctx context.Context, db *sql.DB, userID int, column string) (interface{}, error) {
	var data interface{}
	query := fmt.Sprintf("SELECT %s FROM usuarios WHERE id = $1", column)
	err := db.QueryRowContext(ctx, query, userID).Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func usersSetQuery(ctx context.Context, db *sql.DB, userID int, column string, value interface{}) (bool, error) {

	var query string

	switch value.(type) {
	case int:
		query = fmt.Sprintf("UPDATE usuarios SET %s = %s WHERE id = $1", column, value)
	case string:
		query = fmt.Sprintf("UPDATE usuarios SET %s = '%s' WHERE id = $1", column, value)
	}

	if err := db.QueryRowContext(ctx, query, userID).Err(); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Setters

func SetUserCoinsByID(ctx context.Context, db *sql.DB, userID int, value string) (bool, error) {
	result, err := usersSetQuery(ctx, db, userID, "monedas", value)
	return result, err
}

func SetUserTurnByID(ctx context.Context, db *sql.DB, userID int, value string) (bool, error) {
	result, err := usersSetQuery(ctx, db, userID, "turno", value)
	return result, err
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

func GetUserCoinsByID(ctx context.Context, db *sql.DB, userID int) (int64, error) {
	result, err := usersGetQuery(ctx, db, userID, "monedas")
	return result.(int64), err
}

func GetUserTurnByID(ctx context.Context, db *sql.DB, userID int) (int64, error) {
	result, err := usersGetQuery(ctx, db, userID, "turno")
	return result.(int64), err
}
