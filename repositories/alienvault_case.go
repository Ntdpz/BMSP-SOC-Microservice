package repositories

import (
	"bmsp-backend-service/db"
	"bmsp-backend-service/models"
	"log"
)

func InsertAlienvaultCase(cases []models.AlienvaultCase) error {
	db := db.GetDB()

	if err := db.Create(&cases).Error; err != nil {
		log.Println("Error inserting cases:", err)
		return err
	}

	return nil
}
