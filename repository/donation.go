package repository

import (
	"fmt"
	"holyways/models"

	"gorm.io/gorm"
)

type DonationRepo interface {
	FindDonation() ([]models.Donation, error)
	CreateDonation(donation models.Donation)(models.Donation, error)
	GetDonation(ID int)(models.Donation, error)
	DeleteDonation(donation models.Donation,ID int)(models.Donation, error)
	UpdateDonation(Donation models.Donation) (models.Donation, error)
	GetUserById(ID int) (models.User, error)
	GetFundById(ID int) (models.Fund, error)
	UpdateDonateMidtrans(status string, orderId int) (models.Donation, error)
}
func RepositoryDonation(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindDonation()([]models.Donation, error) {
	var donation []models.Donation

	err := r.db.Preload("User").Preload("Fund").Find(&donation).Error
	return donation , err
}

func (r *repository) CreateDonation(donation models.Donation)(models.Donation, error) {
	err := r.db.Preload("User").Create(&donation).Error
	return donation , err
} 

func (r *repository) GetDonation(ID int) (models.Donation, error) {
	var Donation models.Donation

	err := r.db.Preload("User").Preload("Fund").First(&Donation, ID).Error

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

func (r *repository) GetFundById(ID int) (models.Fund, error) {
	var Fund models.Fund
	err := r.db.Preload("Fund").First(&Fund, ID).Error

	return Fund, err
}

func (r *repository) UpdateDonateMidtrans(status string, orderId int) (models.Donation, error) {
	var donation models.Donation
	r.db.Preload("User").Preload("Fund").First(&donation, orderId)
	
	donation.Status = status
	fmt.Println("repository status : "+ status)
	fmt.Println("repository Doantion status : "+ donation.Status)
	err := r.db.Save(&donation).Error
	return donation, err
}