package repositories

import (
	"bmsp-backend-service/db"
	"bmsp-backend-service/models"
	"log"
)

func InsertNoises(alarms []models.Noises) error {
	db := db.GetDB()

	if err := db.Create(&alarms).Error; err != nil {
		log.Println("Error inserting noises:", err)
		return err
	}

	return nil
}
