package repositories

import (
	"bmsp-backend-service/db"
	"bmsp-backend-service/models"
	"log"
)

// InsertAlienvaultCase ใช้สำหรับบันทึกข้อมูลลงในฐานข้อมูล
func InsertAlienvaultCase(cases []models.AlienvaultCase) error {
	// เชื่อมต่อกับฐานข้อมูล
	db := db.GetDB()

	// ใช้ GORM ในการบันทึกข้อมูลหลายรายการ
	if err := db.Create(&cases).Error; err != nil {
		log.Println("Error inserting cases:", err)
		return err
	}

	return nil
}
