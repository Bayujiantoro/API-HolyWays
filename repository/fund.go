package repository

import (
	"holyways/models"

	"gorm.io/gorm"
)

type FundRepo interface {
	CreateFund(fund models.Fund)(models.Fund, error)
	FindFund()([]models.Fund, error)
	GetFund(ID int)(models.Fund, error)
	UpdateFund(fund models.Fund) (models.Fund, error)
	DeleteFund(fund models.Fund, ID int)(models.Fund, error)

}

func RepositoryFund(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFund()([]models.Fund, error) {
	var fund []models.Fund

	err := r.db.Find(&fund).Error
	return fund , err
}

func (r *repository) CreateFund(fund models.Fund)(models.Fund, error) {
	err := r.db.Create(&fund).Error
	return fund , err
} 

func (r *repository) GetFund(ID int) (models.Fund, error) {
	var fund models.Fund

	err := r.db.First(&fund, ID).Error

	return fund, err
}

func (r *repository) UpdateFund(Fund models.Fund) (models.Fund, error) {
	err := r.db.Save(&Fund).Error

	return Fund, err
}

func (r *repository) DeleteFund(Fund models.Fund, ID int) (models.Fund, error) {
	err := r.db.Delete(&Fund, ID).Scan(&Fund).Error

	return Fund, err
}