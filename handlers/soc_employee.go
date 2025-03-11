package handlers

import (
	"bmsp-backend-service/repositories"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetSocEmployee ดึงข้อมูลพนักงานทั้งหมด
func (h handlers) GetSocEmployee(c *fiber.Ctx) error {
	employees, err := repositories.GetSocEmployees()
	if err != nil {
		log.Println("Error retrieving employees:", err)
		return c.Status(http.StatusInternalServerError).SendString("Failed to retrieve employees")
	}

	return c.JSON(fiber.Map{
		"data":  employees,
		"count": len(employees),
	})
}
