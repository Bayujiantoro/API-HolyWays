package repository

import (
	"holyways/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	CreateUser(models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func(r *repository) FindUsers() ([]models.User, error) {
	var user []models.User
	err := r.db.Find(&user).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {

	err := r.db.Create(&user).Error

	return user, err
}