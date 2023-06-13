package routes

import (
	"holyways/handlers"
	postgre "holyways/pkg/postgresql"
	repository "holyways/repository"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepo := repository.RepositoryUser(postgre.DB)
	h := handlers.HandlerUser(userRepo)

	e.GET("/users", h.FindUsers)
	e.POST("/users", h.CreateUser)
}