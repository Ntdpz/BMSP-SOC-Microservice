package repositories

import (
	"bmsp-backend-service/db"
	"bmsp-backend-service/models"

	"log"
)

// GetSocEmployees ดึงข้อมูลพนักงานทั้งหมดจาก DB และแปลง Time Zone เป็นไทย
func GetSocEmployees() ([]models.SocEmployee, error) {
	var employees []models.SocEmployee
	db := db.GetDB()

	if err := db.Find(&employees).Error; err != nil {
		log.Println("Error retrieving employees:", err)
		return nil, err
	}
	return employees, nil
}

// InsertSocEmployee เพิ่มข้อมูลพนักงานใหม่ในฐานข้อมูล
func InsertSocEmployee(employee models.SocEmployee) error {
	db := db.GetDB()

	// สร้างข้อมูลพนักงานใหม่ในฐานข้อมูล
	if err := db.Create(&employee).Error; err != nil {
		log.Println("Error inserting employee:", err)
		return err
	}

	return nil
}
