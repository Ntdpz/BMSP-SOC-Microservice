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

func GetAlarms(filter models.FilterAlarm) ([]models.Alarm, error) {
	var alarms []models.Alarm
	db := db.GetDB()

	pipeline := db

	if filter.EventStatus != "" {
		pipeline = pipeline.Where("eventstatus = ?", filter.EventStatus) // Change to eventstatus
	}

	if filter.CustomerName != "" {
		pipeline = pipeline.Where("customer_name = ?", filter.CustomerName)
	}

	if err := pipeline.Find(&alarms).Error; err != nil {
		log.Println("Error retrieving alarms with eventstatus filter:", err)
		return nil, err
	}

	return alarms, nil
}

func GetAlarmByID(alarmID string) (*models.Alarm, error) {
	var alarm models.Alarm
	db := db.GetDB()

	if err := db.Where("alarm_id = ?", alarmID).First(&alarm).Error; err != nil {
		log.Println("Error retrieving alarm by alarm_id:", err)
		return nil, err
	}

	return &alarm, nil
}

func UpdateAlarm(alarm *models.Alarm) error {
	db := db.GetDB()

	if err := db.Save(alarm).Error; err != nil {
		log.Println("Error updating alarm:", err)
		return err
	}

	return nil
}
