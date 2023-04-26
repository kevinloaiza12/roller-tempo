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

	name := data.GetAttractionName()
	description := data.GetAttractionDescription()
	duration := data.GetAttractionDuration()
	capacity := data.GetAttractionCapacity()
	currentTurn := data.GetAttractionCurrentTurn()
	nextTurn := data.GetAttractionNextTurn()
	posX := data.GetAttractionPositionX()
	posY := data.GetAttractionPositionY()

	_, err := db.ExecContext(
		ctx,
		"INSERT INTO attractions ("+
			"AttractionName,"+
			"AttractionDescription,"+
			"AttractionDuration,"+
			"AttractionCapacity,"+
			"AttractionCurrentTurn,"+
			"AttractionNextTurn,"+
			"AttractionPosX,"+
			"AttractionPosY)"+
			"VALUES ($1,$2,$3,$4,$5,$6,$7,$8)",
		name,
		description,
		duration,
		capacity,
		currentTurn,
		nextTurn,
		posX,
		posY,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

// Getter

func GetAllAttractions(ctx context.Context, db *sql.DB) ([]map[string]interface{}, error) {

	query := "SELECT * FROM attractions"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var attractions []map[string]interface{}

	for rows.Next() {

		var name string
		var description string
		var capacity int
		var duration int
		var currentTurn int
		var nextTurn int
		var posX float64
		var posY float64

		err := rows.Scan(
			&name,
			&description,
			&capacity,
			&duration,
			&currentTurn,
			&nextTurn,
			&posX,
			&posY,
		)

		if err != nil {
			log.Fatal(err)
		}

		temp := models.NewAttraction(
			name,
			description,
			duration,
			capacity,
			currentTurn,
			nextTurn,
			posX,
			posY,
		)

		attractions = append(attractions, temp.AttractionToJSON())
	}

	return attractions, nil
}

func GetAttractionByName(ctx context.Context, db *sql.DB, attractionName string) (*models.Attraction, error) {

	query := fmt.Sprintf(
		"SELECT %s,%s,%s,%s,%s,%s,%s,%s FROM attractions WHERE AttractionName = $1",
		"AttractionName",
		"AttractionDescription",
		"AttractionDuration",
		"AttractionCapacity",
		"AttractionCurrentTurn",
		"AttractionNextTurn",
		"AttractionPosX",
		"AttractionPosY",
	)

	var name string
	var description string
	var duration int
	var capacity int
	var currentTurn int
	var nextTurn int
	var posX float64
	var posY float64

	err := db.QueryRowContext(ctx, query, attractionName).Scan(
		&name,
		&description,
		&duration,
		&capacity,
		&currentTurn,
		&nextTurn,
		&posX,
		&posY,
	)

	if err != nil {
		return nil, err
	}

	return models.NewAttraction(
		name,
		description,
		duration,
		capacity,
		currentTurn,
		nextTurn,
		posX,
		posY,
	), nil
}

// Update

func AttractionsUpdateQuery(ctx context.Context, db *sql.DB, attraction *models.Attraction) (bool, error) {

	query := fmt.Sprintf(
		"UPDATE attractions SET "+
			"AttractionName = '%s', "+
			"AttractionDescription = '%s', "+
			"AttractionDuration = %d, "+
			"AttractionCapacity = %d, "+
			"AttractionCurrentTurn = %d, "+
			"AttractionNextTurn = %d, "+
			"AttractionPosX = %f, "+
			"AttractionPosY = %f "+
			"WHERE AttractionName = $1",
		attraction.GetAttractionName(),
		attraction.GetAttractionDescription(),
		attraction.GetAttractionDuration(),
		attraction.GetAttractionCapacity(),
		attraction.GetAttractionCurrentTurn(),
		attraction.GetAttractionNextTurn(),
		attraction.GetAttractionPositionX(),
		attraction.GetAttractionPositionY(),
	)

	if _, err := db.ExecContext(ctx, query, attraction.GetAttractionName()); err != nil {
		return false, err
	}

	return true, nil
}

func GetNextTurn(ctx context.Context, db *sql.DB, name string) (int, error) {
	result, err := GetAttractionByName(ctx, db, name)
	if err != nil {
		return 0, err
	}

	currentTurn := result.GetAttractionCurrentTurn()
	nextTurn := currentTurn + 1

	result.SetAttractionCurrentTurn(nextTurn)

	_, updateErr := AttractionsUpdateQuery(ctx, db, result)
	if updateErr != nil {
		return 0, err
	}

	return nextTurn, nil
}
