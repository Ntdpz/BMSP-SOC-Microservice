package handlers

import (
	"bmsp-backend-service/models"
	"bmsp-backend-service/repositories"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (h handlers) CreateNoiseHandler(c *fiber.Ctx) error {
	var requestBody struct {
		Data []models.Noises `json:"data"`
	}

	// Parse the request body
	if err := c.BodyParser(&requestBody); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(400).SendString("Invalid request data")
	}

	// Check if data is provided
	if len(requestBody.Data) == 0 {
		return c.Status(400).SendString("No data provided")
	}

	// Insert data into the database
	err := repositories.InsertNoises(requestBody.Data)
	if err != nil {
		log.Println("Error inserting noise data:", err)
		return c.Status(500).SendString("Failed to insert noise data")
	}

	// Respond with the inserted data
	return c.Status(201).JSON(requestBody.Data)
}
