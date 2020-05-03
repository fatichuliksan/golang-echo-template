package repository

import (
	"project/src/model"

	"github.com/jinzhu/gorm"
)

type userRepo struct {
	DB        *gorm.DB
	TableName string
}

// UserRepoInterface ...
type UserRepoInterface interface {
	FindOneByID(ID uint) (model.User, error)
	FindOneByEmail(email string) (model.User, error)
}

// NewUserRepo ...
func NewUserRepo(db *gorm.DB) UserRepoInterface {
	var model model.User
	return &userRepo{
		DB:        db,
		TableName: model.TableName(),
	}
}

func (t *userRepo) FindOneByID(ID uint) (model.User, error) {
	data := model.User{}
	err := t.DB.First(&data, ID)
	if err.Error != nil {
		return data, err.Error
	}
	return data, nil
}

func (t *userRepo) FindOneByEmail(email string) (model.User, error) {
	data := model.User{}
	err := t.DB.Where("email=$1", email).First(&data)
	if err.Error != nil {
		return data, err.Error
	}
	return data, nil
}
