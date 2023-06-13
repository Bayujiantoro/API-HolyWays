package models


type Donation struct {
	ID   int `json:"id"`
	Date string `json:"date"`
	Money int `json:"money"`
}

type DonationResponse struct {
	ID   int `json:"id"`
	Date string `json:"date"`
	Money int `json:"money"`
}

func(Donation)TableName()string {
	return "Donations"
}