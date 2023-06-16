package funddto

import "holyways/models"

type CreateFund struct {
	Title       string `json:"Title" form:"Title" validate:"required"`
	Description string `json:"Description" from:"Description"`
	Image       string `json:"Image" form:"Image" validate:"required"`
	GoalsMoney  int    `json:"GoalsMoney" form:"GoalsMoney" validate:"required"`
	GoalsDay    string `json:"GoalsDay" form:"GoalsDay" validate:"required"`
	UserID      int    `json:"UserID" form:"UserID" validate:"required"`
}

type FundResponse struct {
	ID          int    `json:"Id"`
	Title       string `json:"Title" `
	Description string `json:"Description"`
	Image       string `json:"Image" form:"image"`
	GoalsMoney  int    `json:"GoalsMoney"`
	GoalsDay    string `json:"GoalsDay"`
	UserID      int    `json:"UserID"`
	User        models.User
	Donation 	[]models.Donation
}