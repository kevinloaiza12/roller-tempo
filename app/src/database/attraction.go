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

// Getter

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

// Update

func AttractionsUpdateQuery(ctx context.Context, db *sql.DB, attraction *resources.Attraction) (bool, error) {

	var query string
	query = fmt.Sprintf(
		"UPDATE atracciones SET nombre = '%s', descripcion = '%s' , duracion = %d, capacidad = %d, siguiente_turno = %d "+
			"WHERE id = $1",
		attraction.GetAttractionName(),
		attraction.GetAttractionDescription(),
		attraction.GetAttractionDuration(),
		attraction.GetAttractionCapacity(),
		attraction.GetAttractionNextTurn(),
	)

	if _, err := db.ExecContext(ctx, query, attraction.GetAttractionID()); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
