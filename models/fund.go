package models

type Fund struct {
	ID          int               `json:"ID" gorm:"primary_key:auto_increment"`
	Title       string            `json:"Title" form:"Title" gorm:"type:varchar(255)"`
	Description string            `json:"Description" gorm:"type:varchar(255)"`
	Image       string            `json:"Image" gorm:"type:varchar(255)"`
	GoalsMoney  int               `json:"GoalsMoney"`
	GoalsDay    string            `json:"GoalsDay"`

	UserID int       `json:"UserID" gorm:"type:int"`
	User   User      `json:"User" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type FundResponse struct {
	ID          int               `json:"ID" gorm:"primary_key:auto_increment"`
	Title       string            `json:"Title" form:"Title" gorm:"type:varchar(255)"`
	Description string            `json:"Description" gorm:"type:varchar(255)"`
	Image       string            `json:"Image" gorm:"type:varchar(255)"`
	GoalsMoney  int               `json:"GoalsMoney"`
	GoalsDay   string            `json:"GoalsDay"`
	UserID int	`json:"UserID"`
	User        UserResponse `json:"User"`
}


func (FundResponse)TableName() string{
	return "funds"
}