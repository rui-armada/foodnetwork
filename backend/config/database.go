package config

import (
	"fmt"
	"log"

	"foodnetwork/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=db user=postgres password=postgres dbname=food port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Post{}, &models.User{}, &models.Business{}, &models.BusinessType{}, &models.BusinessUser{}, &models.JobPost{}, &models.JobTitle{}, &models.Product{}, &models.Service{}, &models.ProfessionalExperience{}, &models.Rating{}, &models.Review{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection established")

	return db
}
