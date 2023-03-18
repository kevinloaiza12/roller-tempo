package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kevinloaiza12/roller-tempo/app/resources"
)

// Creation

func CreateNewReward(ctx context.Context, db *sql.DB, data *resources.Reward) (bool, error) {

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
	} else {
		return true, nil
	}
}

// Getter

func GetRewardByID(ctx context.Context, db *sql.DB, rewardID int) (*resources.Reward, error) {

	query := fmt.Sprintf(
		"SELECT %s,%s,%s,%s FROM premios WHERE id = $1",
		"id",
		"nombre",
		"descripcion",
		"precio",
	)

	var id int64
	var name string
	var description string
	var price int

	err := db.QueryRowContext(ctx, query, rewardID).Scan(
		&id,
		&name,
		&description,
		&price,
	)

	if err != nil {
		return nil, err
	} else {
		return resources.NewReward(
			id,
			name,
			description,
			price,
		), nil
	}
}

// Update

func RewardsUpdateQuery(ctx context.Context, db *sql.DB, reward *resources.Reward) (bool, error) {
	var query string
	query = fmt.Sprintf(
		"UPDATE premios SET nombre = '%s', descripcion = '%s' , precio = %d "+
			"WHERE id = $1",
		reward.GetRewardName(),
		reward.GetRewardDescription(),
		reward.GetRewardPrice(),
	)
	fmt.Println(query)
	if _, err := db.ExecContext(ctx, query, reward.GetRewardID()); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
