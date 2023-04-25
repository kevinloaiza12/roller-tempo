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

	name := data.GetRewardName()
	description := data.GetRewardDescription()
	price := data.GetRewardPrice()

	_, err := db.ExecContext(
		ctx,
		"INSERT INTO rewards (RewardName, RewardDescription, RewardPrice) VALUES ($1,$2,$3)",
		name,
		description,
		price,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

// Getter

func GetAllRewards(ctx context.Context, db *sql.DB) ([]map[string]interface{}, error) {

	query := "SELECT * FROM rewards"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var rewards []map[string]interface{}

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

		rewards = append(rewards, temp.RewardToJSON())
	}

	return rewards, nil
}

func GetRewardByName(ctx context.Context, db *sql.DB, rewardName string) (*models.Reward, error) {

	query := fmt.Sprintf(
		"SELECT %s,%s,%s FROM rewards WHERE RewardName = $1",
		"RewardName",
		"RewardDescription",
		"RewardPrice",
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
		"UPDATE rewards SET RewardName = '%s', RewardDescription = '%s' , RewardPrice = %d "+
			"WHERE RewardName = $1",
		reward.GetRewardName(),
		reward.GetRewardDescription(),
		reward.GetRewardPrice(),
	)

	if _, err := db.ExecContext(ctx, query, reward.GetRewardName()); err != nil {
		return false, err
	}

	return true, nil
}
