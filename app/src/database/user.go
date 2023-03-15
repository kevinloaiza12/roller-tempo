package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func GetUserCoinsByID(ctx context.Context, db *sql.DB, userID int) (int, error) {
	var coins int
	err := db.QueryRowContext(ctx, "SELECT monedas FROM usuarios WHERE id = $1", userID).Scan(&coins)
	if err != nil {
		return 0, err
	}
	return coins, nil
}

func GetUserTurnByID(ctx context.Context, db *sql.DB, userID int) (int, error) {
	var turn int
	err := db.QueryRowContext(ctx, "SELECT turno FROM usuarios WHERE id = $1", userID).Scan(&turn)
	if err != nil {
		return 0, err
	}
	return turn, nil
}
