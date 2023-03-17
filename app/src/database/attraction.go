package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kevinloaiza12/roller-tempo/app/resources"
	_ "github.com/lib/pq"
)

// Creation

func CreateNewAttraction(ctx context.Context, db *sql.DB, data *resources.Attraction) (bool, error) {

	nombre := data.GetAttractionName()
	descripcion := data.GetAttractionDescription()
	duracion := data.GetAttractionDuration()
	capacidad := data.GetAttractionCapacity()
	siguiente_turno := data.GetAttractionNextTurn()

	_, err := db.ExecContext(
		ctx,
		"INSERT INTO atracciones (nombre, descripcion, duracion, capacidad, siguiente_turno) VALUES ($1,$2,$3,$4,$5)",
		nombre,
		descripcion,
		duracion,
		capacidad,
		siguiente_turno,
	)

	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

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

	if _, err := db.ExecContext(ctx, query, attractionID); err != nil {
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

func GetAttractionByID(ctx context.Context, db *sql.DB, attractionID int) (*resources.Attraction, error) {

	query := fmt.Sprintf(
		"SELECT %s,%s,%s,%s,%s,%s FROM atracciones WHERE id = $1",
		"id",
		"nombre",
		"descripcion",
		"duracion",
		"capacidad",
		"siguiente_turno",
	)

	var id int64
	var name string
	var description string
	var duration int
	var capacity int
	var nextTurn int

	err := db.QueryRowContext(ctx, query, attractionID).Scan(
		&id,
		&name,
		&description,
		&duration,
		&capacity,
		&nextTurn,
	)

	if err != nil {
		return nil, err
	} else {
		return resources.NewAttraction(
			id,
			name,
			description,
			duration,
			capacity,
			nextTurn,
		), nil
	}
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
