package main

import (
	"fmt"
	"holyways/database"
	postgre "holyways/pkg/postgresql"
	"holyways/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	 godotenv.Load()


	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	  }))

	postgre.DatabaseConnection()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))
	var PORT = os.Getenv("PORT")

	fmt.Println("Running on Port : 5200")
	e.Logger.Fatal(e.Start(":" + PORT))
}
