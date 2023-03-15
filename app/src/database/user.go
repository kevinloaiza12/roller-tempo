package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func usersGetQuery(ctx context.Context, db *sql.DB, attractionID int, column string) (interface{}, error) {
	var data interface{}
	query := fmt.Sprintf("SELECT %s FROM usuarios WHERE id = $1", column)
	err := db.QueryRowContext(ctx, query, attractionID).Scan(&data)
	if err != nil {
		return 0, err
	}
	return data, nil
}

func GetUserCoinsByID(ctx context.Context, db *sql.DB, attractionID int) (int64, error) {
	result, err := usersGetQuery(ctx, db, attractionID, "monedas")
	return result.(int64), err
}

func GetUserTurnByID(ctx context.Context, db *sql.DB, attractionID int) (int64, error) {
	result, err := usersGetQuery(ctx, db, attractionID, "turno")
	return result.(int64), err
}
