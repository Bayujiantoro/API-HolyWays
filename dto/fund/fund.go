package funddto


type CreateFund struct {
	Title       string `json:"Title" form:"Title" validate:"required"`
	Description string `json:"Description" from:"Description"`
	Image string `json:"Image" form:"Image" validate:"required"`
	Goals_money int `json:"Goals_money" form:"Goals_money" validate:"required"`
	Goals_day   string `json:"Goals_day" form:"Goals_day" validate:"required"`
}

type FundResponse struct {
	ID int `json:"id"`
	Title       string `json:"title" `
	Description string `json:"description"`
	Image string `json:"image" form:"image"`
	Goals_money int `json:"goals_money"`
	Goals_day   string `json:"goals_day"`
}