package models

import "time"

type Donation struct {
	ID   int `json:"id"`
	Date time.Time `json:"date"`
	Money int `json:"money"`
}

type DonationResponse struct {
	ID   int `json:"id"`
	Date time.Time `json:"date"`
	Money int `json:"money"`
}

func(Donation)TableName()string {
	return "Donations"
}