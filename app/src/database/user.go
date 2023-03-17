package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	id    int
	coins int
	turn  int
}

// Creation

func NewUser(id int, coins int, turn int) *User {
	return &User{
		id,
		coins,
		turn,
	}
}

func CreateNewUser(ctx context.Context, db *sql.DB, data *User) (bool, error) {

	id := data.id
	monedas := data.coins
	turno := data.turn

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

func GetUserByID(ctx context.Context, db *sql.DB, userID int) (*User, error) {

	var user User

	query := fmt.Sprintf(
		"SELECT %s,%s,%s FROM usuarios WHERE id = $1",
		"id",
		"monedas",
		"turno",
	)

	err := db.QueryRowContext(ctx, query, userID).Scan(
		&user.id,
		&user.coins,
		&user.turn,
	)

	if err != nil {
		return nil, err
	} else {
		return &user, nil
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
