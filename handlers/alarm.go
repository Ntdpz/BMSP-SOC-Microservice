package handlers

import (
	"bmsp-backend-service/models"
	"bmsp-backend-service/repositories"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h handlers) CreateAlarmHandler(c *fiber.Ctx) error {
	var requestBody struct {
		Data []models.Alarm `json:"data"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(400).SendString("Invalid request data")
	}

	if len(requestBody.Data) == 0 {
		return c.Status(400).SendString("No data provided")
	}

	err := repositories.InsertAlarm(requestBody.Data)
	if err != nil {
		log.Println("Error inserting alarm data:", err)
		return c.Status(500).SendString("Failed to insert alarms")
	}

	return c.Status(201).JSON(requestBody.Data)
}
func (h handlers) GetAlarms(c *fiber.Ctx) error {
	isOpenStr := c.Query("is_open")
	var isOpen *bool

	if isOpenStr != "" {
		parsedIsOpen, err := strconv.ParseBool(isOpenStr)
		if err != nil {
			log.Println("Invalid is_open value:", err)
			return c.Status(400).SendString("Invalid is_open value")
		}
		isOpen = &parsedIsOpen
	}

	alarms, err := repositories.GetAlarms(isOpen)
	if err != nil {
		log.Println("Error retrieving alarms:", err)
		return c.Status(500).SendString("Failed to retrieve alarms")
	}

	return c.JSON(fiber.Map{
		"data":  alarms,
		"count": len(alarms),
	})
}
