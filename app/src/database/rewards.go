package database

import (
	"context"
	"database/sql"
	"fmt"
)

type Reward struct {
	id          int64
	name        string
	description string
	price       int
}

// Creation

func NewReward(id int64, name string, description string, price int) *Reward {
	return &Reward{
		id,
		name,
		description,
		price,
	}
}

func CreateNewReward(ctx context.Context, db *sql.DB, data *Reward) (bool, error) {

	nombre := data.name
	descripcion := data.description
	precio := data.price

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

// Utils

func rewardsGetQuery(ctx context.Context, db *sql.DB, attractionID int, column string) (interface{}, error) {
	var data interface{}
	query := fmt.Sprintf("SELECT %s FROM premios WHERE id = $1", column)
	err := db.QueryRowContext(ctx, query, attractionID).Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func rewardsSetQuery(ctx context.Context, db *sql.DB, attractionID int, column string, value interface{}) (bool, error) {

	var query string

	switch value.(type) {
	case int:
		query = fmt.Sprintf("UPDATE premios SET %s = %s WHERE id = $1", column, value)
	case string:
		query = fmt.Sprintf("UPDATE premios SET %s = '%s' WHERE id = $1", column, value)
	}

	if err := db.QueryRowContext(ctx, query, attractionID).Err(); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Setters

func SetRewardNameByID(ctx context.Context, db *sql.DB, attractionID int, value string) (bool, error) {
	result, err := rewardsSetQuery(ctx, db, attractionID, "nombre", value)
	return result, err
}

func SetRewardDescriptionByID(ctx context.Context, db *sql.DB, attractionID int, value string) (bool, error) {
	result, err := rewardsSetQuery(ctx, db, attractionID, "descripcion", value)
	return result, err
}

func SetRewardPriceByID(ctx context.Context, db *sql.DB, attractionID int, value int) (bool, error) {
	result, err := rewardsSetQuery(ctx, db, attractionID, "precio", value)
	return result, err
}

// Getters

func GetRewardByID(ctx context.Context, db *sql.DB, rewardID int) (*Reward, error) {

	var reward Reward

	query := fmt.Sprintf(
		"SELECT %s,%s,%s,%s FROM premios WHERE id = $1",
		"id",
		"nombre",
		"descripcion",
		"precio",
	)

	err := db.QueryRowContext(ctx, query, rewardID).Scan(
		&reward.id,
		&reward.name,
		&reward.description,
		&reward.price,
	)

	if err != nil {
		return nil, err
	} else {
		return &reward, nil
	}
}

func GetRewardNameByID(ctx context.Context, db *sql.DB, rewardID int) (string, error) {
	result, err := rewardsGetQuery(ctx, db, rewardID, "nombre")
	return result.(string), err
}

func GetRewardDescriptionByID(ctx context.Context, db *sql.DB, rewardID int) (string, error) {
	result, err := rewardsGetQuery(ctx, db, rewardID, "descripcion")
	return result.(string), err
}

func GetRewardPriceByID(ctx context.Context, db *sql.DB, rewardID int) (int, error) {
	result, err := rewardsGetQuery(ctx, db, rewardID, "precio")
	return result.(int), err
}
