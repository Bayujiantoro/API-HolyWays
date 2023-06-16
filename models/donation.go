package models

import "time"

type Donation struct {
	ID    int       `json:"ID" gorm:"primary_key:auto_increment"`
	Date  time.Time `json:"Date"`
	Money int       `json:"Money"`
	Status string `json:"Status" gorm:"type:varchar(200)"`
	FundID int    `json:"FundID" gorm:"type: int"`
	Fund   Fund   `json:"Fund" gorm:"foreignKey:FundID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	UserID int  `json:"UserID" gorm:"type: int"`
	User   User	`json:"User" gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type DonationResponse struct {
	ID     int       `json:"id"`
	Date   time.Time `json:"date"`
	Status      string       `json:"Status" gorm:"type:varchar(200)"`
	Money  int       `json:"money"`
	FundID int       `json:"FundID"`

	UserID int          `json:"UserID"`
}

func (Donation) TableName() string {
	return "donations"
}
