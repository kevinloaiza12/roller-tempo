package database

import (
	"context"
	"database/sql"
	"fmt"
)

func rewardsGetQuery(ctx context.Context, db *sql.DB, attractionID int, column string) (interface{}, error) {
	var data interface{}
	query := fmt.Sprintf("SELECT %s FROM premios WHERE id = $1", column)
	err := db.QueryRowContext(ctx, query, attractionID).Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetRewardsNameByID(ctx context.Context, db *sql.DB, attractionID int) (string, error) {
	result, err := rewardsGetQuery(ctx, db, attractionID, "nombre")
	return result.(string), err
}

func GetRewardDescriptionByID(ctx context.Context, db *sql.DB, attractionID int) (string, error) {
	result, err := rewardsGetQuery(ctx, db, attractionID, "descripcion")
	return result.(string), err
}

func GetRewardPriceByID(ctx context.Context, db *sql.DB, attractionID int) (int, error) {
	result, err := rewardsGetQuery(ctx, db, attractionID, "precio")
	return result.(int), err
}
