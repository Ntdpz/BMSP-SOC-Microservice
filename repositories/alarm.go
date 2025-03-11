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

func GetAlarms(isOpen *bool) ([]models.Alarm, error) {
	var alarms []models.Alarm
	db := db.GetDB()

	if isOpen != nil {
		if err := db.Where("is_open = ?", *isOpen).Find(&alarms).Error; err != nil {
			log.Println("Error retrieving alarms with is_open filter:", err)
			return nil, err
		}
	} else {
		if err := db.Find(&alarms).Error; err != nil {
			log.Println("Error retrieving alarms:", err)
			return nil, err
		}
	}

	return alarms, nil
}
