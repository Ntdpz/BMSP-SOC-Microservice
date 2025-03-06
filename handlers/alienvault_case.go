package handlers

import (
	"bmsp-backend-service/models"
	"bmsp-backend-service/repositories"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateAlienvaultCaseHandler(c *fiber.Ctx) error {
	var requestBody struct {
		Data []models.AlienvaultCase `json:"data"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(400).SendString("Invalid request data")
	}

	if len(requestBody.Data) == 0 {
		return c.Status(400).SendString("No data provided")
	}

	err := repositories.InsertAlienvaultCase(requestBody.Data)
	if err != nil {
		return c.Status(500).SendString("Failed to insert cases")
	}

	return c.Status(201).JSON(requestBody.Data)
}
