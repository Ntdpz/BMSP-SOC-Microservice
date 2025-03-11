package handlers

import (
	"bmsp-backend-service/models"
	"bmsp-backend-service/repositories"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h handlers) CreateNoiseHandler(c *fiber.Ctx) error {
	var requestBody struct {
		Data []models.Noises `json:"data"`
	}

	// Parse the request body
	if err := c.BodyParser(&requestBody); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(http.StatusBadRequest).SendString("Invalid request data")
	}

	// Check if data is provided
	if len(requestBody.Data) == 0 {
		return c.Status(http.StatusBadRequest).SendString("No data provided")
	}

	// Insert data into the database
	err := repositories.InsertNoises(requestBody.Data)
	if err != nil {
		log.Println("Error inserting noise data:", err)
		return c.Status(http.StatusInternalServerError).SendString("Failed to insert noise data")
	}

	// Respond with the inserted data
	return c.Status(http.StatusCreated).JSON(requestBody.Data)
}
