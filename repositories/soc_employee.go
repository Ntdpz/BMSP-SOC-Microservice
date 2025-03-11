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

	// แปลงเวลาของพนักงานทุกคนเป็นโซนไทย
	// for i := range employees {
	// 	employees[i].CreatedAt = utils.ConvertToThaiTime(employees[i].CreatedAt)
	// 	employees[i].UpdatedAt = utils.ConvertToThaiTime(employees[i].UpdatedAt)
	// }

	return employees, nil
}
