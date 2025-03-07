package repositories

import (
	"bmsp-backend-service/db"
	"bmsp-backend-service/models"
	"log"
)

func InsertAlarm(alarms []models.Alarm) error {
	db := db.GetDB()

	if err := db.Create(&alarms).Error; err != nil {
		log.Println("Error inserting alarms:", err)
		return err
	}

	return nil
}
