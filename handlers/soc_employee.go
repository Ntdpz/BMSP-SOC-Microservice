package handlers

import (
	"bmsp-backend-service/models"
	"bmsp-backend-service/repositories"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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

func (h handlers) CreateSocEmployee(c *fiber.Ctx) error {
	// รับข้อมูลจาก body
	var requestBody models.SocEmployee

	if requestBody.IsActive == nil {
		defaultIsActive := true
		requestBody.IsActive = &defaultIsActive
	}

	if err := c.BodyParser(&requestBody); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(http.StatusBadRequest).SendString("Invalid request data")
	}

	requestBody.Key = uuid.New().String()

	if requestBody.FirstName == "" || requestBody.LastName == "" {
		return c.Status(http.StatusBadRequest).SendString("First name and last name are required")
	}

	err := repositories.InsertSocEmployee(requestBody)
	if err != nil {
		log.Println("Error inserting employee data:", err)
		return c.Status(http.StatusInternalServerError).SendString("Failed to insert employee")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Employee created successfully",
		// "data":    requestBody,
	})
}
