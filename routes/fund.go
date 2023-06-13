package routes

import (
	"holyways/handlers"
	"holyways/pkg/middleware"
	postgre "holyways/pkg/postgresql"
	"holyways/repository"

	"github.com/labstack/echo/v4"
)

func FundRoutes(e *echo.Group) {
	fund := repository.RepositoryFund(postgre.DB)

	h := handlers.FundHandler(fund)

	e.POST("/fund", middleware.UploadFile(h.CreateFund))
	e.GET("/fund", h.FindFund)
	e.GET("/fund/:id", h.GetFund)
	e.PATCH("/fund/:id", middleware.UploadFile(h.UpdateFund))
	e.DELETE("/fund/:id", h.DeleteFund)
}