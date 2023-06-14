package userdto

import "holyways/models"

type CreateUser struct {
	Name     string `json:"Name" form:"Name" validate:"required"`
	Email    string `json:"Email" form:"Email" validate:"required"`
	Phone    string `json:"Phone" form:"Phone" validate:"required"`
	Password string `json:"Password" form:"Password" validate:"required"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"Phone"`
	Fund	[]models.Fund	
	Donation []models.Donation
}
