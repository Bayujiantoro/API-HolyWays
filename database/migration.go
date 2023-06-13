package database

import (
	"fmt"
	"holyways/models"
	postgre "holyways/pkg/postgresql"
)

// automatic migrate identifikasi

func RunMigration() {
	err := postgre.DB.AutoMigrate(
		&models.User{},
		&models.Fund{},
		&models.Donation{},
	)
	if err !=nil {
		panic("Migration Failed")
	}

	fmt.Println("Migration Succes")
}