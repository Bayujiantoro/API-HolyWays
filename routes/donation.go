package routes

import (
	"holyways/handlers"
	postgre "holyways/pkg/postgresql"
	"holyways/repository"

	"github.com/labstack/echo/v4"
)

func DonationRoutes(e *echo.Group) {
	donation := repository.RepositoryDonation(postgre.DB)
	h := handlers.DonationHandler(donation)

	e.POST("/donation", h.CreateDonation)
	e.GET("/donation", h.FindDonation)
	e.GET("/donation/:id", h.GetDonation)
	e.PATCH("/donation/:id", h.UpdateDonation)
	e.DELETE("/donation:id", h.DeleteDonation)
}