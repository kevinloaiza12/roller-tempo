package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/kevinloaiza12/roller-tempo/app/models"
	_ "github.com/lib/pq"
)

// Creation

func CreateNewAttraction(ctx context.Context, db *sql.DB, data *models.Attraction) (bool, error) {

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
	}

	return true, nil
}

// Getter

func GetAllAttractions(ctx context.Context, db *sql.DB) ([]map[string]interface{}, error) {

	query := "SELECT * FROM atracciones"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var atracciones []map[string]interface{}

	for rows.Next() {

		var name string
		var description string
		var duration int
		var capacity int
		var nextTurn int

		err := rows.Scan(
			&name,
			&description,
			&duration,
			&capacity,
			&nextTurn,
		)

		if err != nil {
			log.Fatal(err)
		}

		temp := models.NewAttraction(
			name,
			description,
			duration,
			capacity,
			nextTurn,
		)

		atracciones = append(atracciones, temp.AttractionToJSON())
	}

	return atracciones, nil
}

func GetAttractionByName(ctx context.Context, db *sql.DB, attractionName string) (*models.Attraction, error) {

	query := fmt.Sprintf(
		"SELECT %s,%s,%s,%s,%s FROM atracciones WHERE nombre = $1",
		"nombre",
		"descripcion",
		"duracion",
		"capacidad",
		"siguiente_turno",
	)

	var name string
	var description string
	var duration int
	var capacity int
	var nextTurn int

	err := db.QueryRowContext(ctx, query, attractionName).Scan(
		&name,
		&description,
		&duration,
		&capacity,
		&nextTurn,
	)

	if err != nil {
		return nil, err
	}

	return models.NewAttraction(
		name,
		description,
		duration,
		capacity,
		nextTurn,
	), nil
}

// Update

func AttractionsUpdateQuery(ctx context.Context, db *sql.DB, attraction *models.Attraction) (bool, error) {

	query := fmt.Sprintf(
		"UPDATE atracciones SET nombre = '%s', descripcion = '%s' , duracion = %d, capacidad = %d, siguiente_turno = %d "+
			"WHERE nombre = $1",
		attraction.GetAttractionName(),
		attraction.GetAttractionDescription(),
		attraction.GetAttractionDuration(),
		attraction.GetAttractionCapacity(),
		attraction.GetAttractionNextTurn(),
	)

	if _, err := db.ExecContext(ctx, query, attraction.GetAttractionName()); err != nil {
		return false, err
	}

	return true, nil
}
