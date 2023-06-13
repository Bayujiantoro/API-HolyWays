package donationdto

type CreateDonation struct {
	Date string `json:"Date" validate:"required"`
	Money int `json:"Money" validate:"required"`
}

type DonationResponse struct {
	ID   int `json:"ID"`
	Date string `json:"Date"`
	Money int `json:"Money"`
}