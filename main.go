package main

import (
	"fmt"
	"holyways/database"
	postgre "holyways/pkg/postgresql"
	"holyways/routes"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("env failed")
	}

	e := echo.New()

	postgre.DatabaseConnection()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	fmt.Println("Running on Port : 5200")
	e.Logger.Fatal(e.Start("localhost:5200"))
}

func helloworld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}