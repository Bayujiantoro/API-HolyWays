package repository

import (
	"holyways/models"

	"gorm.io/gorm"
)

type DonationRepo interface {
	FindDonation() ([]models.Donation, error)
	CreateDonation(donation models.Donation)(models.Donation, error)
	GetDonation(ID int)(models.Donation, error)
	DeleteDonation(donation models.Donation,ID int)(models.Donation, error)
	UpdateDonation(Donation models.Donation) (models.Donation, error)
}
func RepositoryDonation(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindDonation()([]models.Donation, error) {
	var donation []models.Donation

	err := r.db.Find(&donation).Error
	return donation , err
}

func (r *repository) CreateDonation(donation models.Donation)(models.Donation, error) {
	err := r.db.Create(&donation).Error
	return donation , err
} 

func (r *repository) GetDonation(ID int) (models.Donation, error) {
	var Donation models.Donation

	err := r.db.First(&Donation, ID).Error

	return Donation, err
}

func (r *repository) UpdateDonation(Donation models.Donation) (models.Donation, error) {
	err := r.db.Save(&Donation).Error

	return Donation, err
}

func (r *repository) DeleteDonation(Donation models.Donation, ID int) (models.Donation, error) {
	err := r.db.Delete(&Donation, ID).Scan(&Donation).Error

	return Donation, err
}