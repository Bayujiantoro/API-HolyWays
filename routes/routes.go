package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	FundRoutes(e)
	DonationRoutes(e)
	AuthRoutes(e)
}