package routes

import (
	"holyways/handlers"
	"holyways/pkg/middleware"
	postgre "holyways/pkg/postgresql"
	repository "holyways/repository"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepo := repository.RepositoryUser(postgre.DB)
	h := handlers.HandlerUser(userRepo)

	e.GET("/users", h.FindUsers)
	e.POST("/users", middleware.Auth(h.CreateUser))
	e.GET("/users/:id", h.GetUser)
	e.PATCH("/users/:id", middleware.Auth(h.UpdateUser))
	e.DELETE("/users/:id", middleware.Auth(h.DeleteUser))
}