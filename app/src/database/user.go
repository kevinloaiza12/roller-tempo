package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func usersQuery(ctx context.Context, db *sql.DB, attractionID int, column string) (interface{}, error) {
	var data interface{}
	err := db.QueryRowContext(ctx, "SELECT $1 FROM usuarios WHERE id = $2", column, attractionID).Scan(&data)
	if err != nil {
		return 0, err
	}
	return data, nil
}

func GetUserCoinsByID(ctx context.Context, db *sql.DB, attractionID int) (int, error) {
	result, err := usersQuery(ctx, db, attractionID, "monedas")
	return result.(int), err
}

func GetUserTurnByID(ctx context.Context, db *sql.DB, attractionID int) (int, error) {
	result, err := usersQuery(ctx, db, attractionID, "turno")
	return result.(int), err
}
