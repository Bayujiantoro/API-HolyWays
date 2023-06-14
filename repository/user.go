package repository

import (
	"holyways/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUser() ([]models.User, error)
	CreateUser(models.User) (models.User, error)
	GetUser(ID int)(models.User, error)
	UpdateUser(User models.User) (models.User, error)
	DeleteUser(User models.User, ID int)(models.User, error)
	
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func(r *repository) FindUser() ([]models.User, error) {
	var user []models.User
	err := r.db.Find(&user).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {

	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var User models.User

	err := r.db.First(&User, ID).Error

	return User, err
}

func (r *repository) UpdateUser(User models.User) (models.User, error) {
	err := r.db.Save(&User).Error

	return User, err
}

func (r *repository) DeleteUser(User models.User, ID int) (models.User, error) {
	err := r.db.Delete(&User, ID).Scan(&User).Error

	return User, err
}