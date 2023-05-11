package repository

import (
	"roller-tempo/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) UpdateUser(user *model.User) error {
	return ur.db.Save(user).Error
}

func (ur *UserRepository) DeleteUser(user *model.User) error {
	return ur.db.Delete(user).Error
}

func (ur *UserRepository) GetUserByID(id int) (*model.User, error) {
	var user model.User
	err := ur.db.Where("identification = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
