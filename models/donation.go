package models

import "time"

type Donation struct {
	ID    int       `json:"ID" gorm:"primary_key:auto_increment"`
	Date  time.Time `json:"Date"`
	Money int       `json:"Money"`

	FundID int `json:"FundID"`
	Fund Fund `json:"Fund"`
	UserID int  `json:"UserID" gorm:"type: int"`
	User   User	`json:"User" gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type DonationResponse struct {
	ID     int       `json:"id"`
	Date   time.Time `json:"date"`
	Money  int       `json:"money"`
	FundID int       `json:"FundID"`

	UserID int          `json:"UserID"`
	User   UserResponse `json:"User"`
}

func (Donation) TableName() string {
	return "Donations"
}
