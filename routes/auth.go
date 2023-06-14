package routes

import (
	"holyways/handlers"
	"holyways/pkg/middleware"
	postgre "holyways/pkg/postgresql"
	"holyways/repository"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	authRepo := repository.RepositoryAuth(postgre.DB)

	h := handlers.HandlerAuth(authRepo)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))
}