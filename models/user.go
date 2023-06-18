package models

type User struct {
	ID       int    `json:"ID" gorm:"primary_key:auto_increment"`
	Name     string `json:"Name" gorm:"type:varchar(255)"`
	Email    string `json:"Email" gorm:"type:varchar(255)"`
	Password string `json:"Password" gorm:"type:varchar(255)"`
	Phone    string `json:"Phone"`
	ProfilePicture string	`json:"ProfilePicture" gorm:"type:varchar(255)"`
	Fund     []Fund `json:"Funds" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Donation []Donation `json:"Donation" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}


type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Phone string `json:"phone"`
	ProfilePicture string	`json:"ProfilePicture" gorm:"type:varchar(255)"`
	Fund Fund `json:"Fund"`
	Donation Donation	`json:"Donation"`
}

func (UserResponse) TableName() string {
	return "users"
}