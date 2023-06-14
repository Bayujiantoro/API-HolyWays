package routes

import (
	"holyways/handlers"
	"holyways/pkg/middleware"
	postgre "holyways/pkg/postgresql"
	"holyways/repository"

	"github.com/labstack/echo/v4"
)

func DonationRoutes(e *echo.Group) {
	donation := repository.RepositoryDonation(postgre.DB)
	h := handlers.DonationHandler(donation)

	e.POST("/donation", middleware.Auth(h.CreateDonation))
	e.GET("/donation", middleware.Auth(h.FindDonation))
	e.GET("/donation/:id", middleware.Auth(h.GetDonation))
	e.PATCH("/donation/:id", middleware.Auth(h.UpdateDonation))
	e.DELETE("/donation/:id", middleware.Auth(h.DeleteDonation))
}