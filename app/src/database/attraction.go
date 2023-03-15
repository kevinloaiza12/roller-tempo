package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func GetNextTurnByID(ctx context.Context, db *sql.DB, attractionID int) (int, error) {
	var turn int
	err := db.QueryRowContext(ctx, "SELECT siguiente_turno FROM usuarios WHERE id = $1", attractionID).Scan(&turn)
	if err != nil {
		return 0, err
	}
	return turn, nil
}
