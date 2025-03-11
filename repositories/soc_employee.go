package repositories

import (
	"bmsp-backend-service/db"
	"bmsp-backend-service/models"

	"log"
)

func GetSocEmployees() ([]models.SocEmployee, error) {
	var employees []models.SocEmployee
	db := db.GetDB()

	if err := db.Find(&employees).Error; err != nil {
		log.Println("Error retrieving employees:", err)
		return nil, err
	}
	return employees, nil
}

func InsertSocEmployee(employee models.SocEmployee) error {
	db := db.GetDB()

	if err := db.Create(&employee).Error; err != nil {
		log.Println("Error inserting employee:", err)
		return err
	}

	return nil
}
func GetSocEmployeeByID(id string) (models.SocEmployee, error) {
	var employee models.SocEmployee
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&employee).Error; err != nil {
		log.Println("Error retrieving employee by ID:", err)
		return models.SocEmployee{}, err
	}

	return employee, nil
}
func UpdateSocEmployee(id string, employee models.SocEmployee) error {
	db := db.GetDB()

	if err := db.Model(&models.SocEmployee{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"first_name": employee.FirstName,
			"last_name":  employee.LastName,
			"is_active":  employee.IsActive,
			"updated_at": employee.UpdatedAt,
		}).Error; err != nil {
		log.Println("Error updating employee:", err)
		return err
	}

	return nil
}
