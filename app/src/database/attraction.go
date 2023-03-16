package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Utils

func attractionsGetQuery(ctx context.Context, db *sql.DB, attractionID int, column string) (interface{}, error) {
	var data interface{}
	query := fmt.Sprintf("SELECT %s FROM atracciones WHERE id = $1", column)
	err := db.QueryRowContext(ctx, query, attractionID).Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func attractionsSetQuery(ctx context.Context, db *sql.DB, attractionID int, column string, value interface{}) (bool, error) {

	var query string

	switch value.(type) {
	case int:
		query = fmt.Sprintf("UPDATE atracciones SET %s = %s WHERE id = $1", column, value)
	case string:
		query = fmt.Sprintf("UPDATE atracciones SET %s = '%s' WHERE id = $1", column, value)
	}

	if err := db.QueryRowContext(ctx, query, attractionID).Err(); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Setters

func SetAttractionNameByID(ctx context.Context, db *sql.DB, attractionID int, value string) (bool, error) {
	result, err := attractionsSetQuery(ctx, db, attractionID, "nombre", value)
	return result, err
}

func SetAttractionDescriptionByID(ctx context.Context, db *sql.DB, attractionID int, value string) (bool, error) {
	result, err := attractionsSetQuery(ctx, db, attractionID, "descripcion", value)
	return result, err
}

func SetAttractionDurationByID(ctx context.Context, db *sql.DB, attractionID int, value int) (bool, error) {
	result, err := attractionsSetQuery(ctx, db, attractionID, "duracion", value)
	return result, err
}

func SetAttractionCapacityByID(ctx context.Context, db *sql.DB, attractionID int, value int) (bool, error) {
	result, err := attractionsSetQuery(ctx, db, attractionID, "capacidad", value)
	return result, err
}

func SetAttractionNextTurnByID(ctx context.Context, db *sql.DB, attractionID int, value int) (bool, error) {
	result, err := attractionsSetQuery(ctx, db, attractionID, "siguiente_turno", value)
	return result, err
}

// Getters

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
