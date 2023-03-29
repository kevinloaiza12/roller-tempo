package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/kevinloaiza12/roller-tempo/app/models"
)

// Creation

func CreateNewReward(ctx context.Context, db *sql.DB, data *models.Reward) (bool, error) {

	nombre := data.GetRewardName()
	descripcion := data.GetRewardDescription()
	precio := data.GetRewardPrice()

	_, err := db.ExecContext(
		ctx,
		"INSERT INTO premios (nombre, descripcion, precio) VALUES ($1,$2,$3)",
		nombre,
		descripcion,
		precio,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

// Getter

func GetAllRewards(ctx context.Context, db *sql.DB) ([]map[string]interface{}, error) {

	query := "SELECT * FROM premios"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var premios []map[string]interface{}

	for rows.Next() {

		var name string
		var description string
		var price int

		err := rows.Scan(
			&name,
			&description,
			&price,
		)

		if err != nil {
			log.Fatal(err)
		}

		temp := models.NewReward(
			name,
			description,
			price,
		)

		premios = append(premios, temp.RewardToJSON())
	}

	return premios, nil
}

func GetRewardByName(ctx context.Context, db *sql.DB, rewardName string) (*models.Reward, error) {

	query := fmt.Sprintf(
		"SELECT %s,%s,%s FROM premios WHERE nombre = $1",
		"nombre",
		"descripcion",
		"precio",
	)

	var name string
	var description string
	var price int

	err := db.QueryRowContext(ctx, query, rewardName).Scan(
		&name,
		&description,
		&price,
	)

	if err != nil {
		return nil, err
	}

	return models.NewReward(
		name,
		description,
		price,
	), nil
}

// Update

func RewardsUpdateQuery(ctx context.Context, db *sql.DB, reward *models.Reward) (bool, error) {

	query := fmt.Sprintf(
		"UPDATE premios SET nombre = '%s', descripcion = '%s' , precio = %d "+
			"WHERE nombre = $1",
		reward.GetRewardName(),
		reward.GetRewardDescription(),
		reward.GetRewardPrice(),
	)

	if _, err := db.ExecContext(ctx, query, reward.GetRewardName()); err != nil {
		return false, err
	}

	return true, nil
}
