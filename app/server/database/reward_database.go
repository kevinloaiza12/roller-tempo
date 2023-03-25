package database

import (
	"context"
	"database/sql"
	"fmt"

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
