package models

type Fund struct {
	ID int `json:"ID" gorm:"primary_key:auto_increment"`
	Title      string          `json:"Title" form:"Title" gorm:"type: varchar(255)"`
	Description string `json:"Description" gorm:"type: varchar(255)"`
	Image string `json:"Image" gorm:"type: varchar(255)"`
	Goals_money int `json:"Goals_money"`
	Goals_day string `json:"Goals_day"`
}

type FundResponse struct {
	ID int `json:"ID" gorm:"primary_key:auto_increment"`
	Title      string          `json:"Title" form:"Title" gorm:"type: varchar(255)"`
	Description string `json:"Description" gorm:"type: varchar(255)"`
	Image string `json:"Image" gorm:"type: varchar(255)"`
	Goals_money int `json:"Goals_money"`
	Goals_day string `json:"Goals_day"`
}

func (FundResponse)TableName() string{
	return "funds"
}