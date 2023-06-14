package donationdto

import (
	"holyways/models"
	"time"
)

type CreateDonation struct {
	Date   time.Time `json:"Date" validate:"required"`
	Money  int       `json:"Money" validate:"required"`
	UserID int       `json:"UserID" validate:"required"`
	FundID int       `json:"FundID" validate:"required"`
}

type DonationResponse struct {
	ID     int       `json:"ID"`
	Date   time.Time `json:"Date"`
	Money  int       `json:"Money"`
	UserID int       `json:"UserID"`
	User   models.User
	FundID int `json:"FundID"`
	Fund   models.Fund
}
