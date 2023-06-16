package donationdto

import (
	"holyways/models"
	"time"
)

type CreateDonation struct {
	Date   time.Time `json:"Date"`
	Money  int       `json:"Money"`
	Status string    `json:"Status"`
	UserID int       `json:"UserID"`
	FundID int       `json:"FundID" validate:"required"`
}

type DonationResponse struct {
	ID     int       `json:"ID"`
	Date   time.Time `json:"Date"`
	Money  int       `json:"Money"`
	Status string    `json:"Status"`
	UserID int       `json:"UserID"`
	User   models.User
	FundID int `json:"FundID"`
	Fund   models.Fund
}
