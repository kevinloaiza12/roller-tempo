package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func attractionsGetQuery(ctx context.Context, db *sql.DB, attractionID int, column string) (interface{}, error) {
	var data interface{}
	query := fmt.Sprintf("SELECT %s FROM atracciones WHERE id = $1", column)
	err := db.QueryRowContext(ctx, query, attractionID).Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetAttractionNameByID(ctx context.Context, db *sql.DB, attractionID int) (string, error) {
	result, err := attractionsGetQuery(ctx, db, attractionID, "nombre")
	return result.(string), err
}

func GetAttractionDescriptionByID(ctx context.Context, db *sql.DB, attractionID int) (string, error) {
	result, err := attractionsGetQuery(ctx, db, attractionID, "descripcion")
	return result.(string), err
}

func GetAttractionDurationByID(ctx context.Context, db *sql.DB, attractionID int) (int, error) {
	result, err := attractionsGetQuery(ctx, db, attractionID, "duracion")
	return result.(int), err
}

func GetAttractionCapacityByID(ctx context.Context, db *sql.DB, attractionID int) (int, error) {
	result, err := attractionsGetQuery(ctx, db, attractionID, "capacidad")
	return result.(int), err
}

func GetAttractionNextTurnByID(ctx context.Context, db *sql.DB, attractionID int) (int, error) {
	result, err := attractionsGetQuery(ctx, db, attractionID, "siguiente_turno")
	return result.(int), err
}
