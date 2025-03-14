package handlers

import (
	"bmsp-backend-service/models"
	"bmsp-backend-service/repositories"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h handlers) CreateAlarmHandler(c *fiber.Ctx) error {
	var requestBody struct {
		Data []models.Alarm `json:"data"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(http.StatusBadRequest).SendString("Invalid request data")
	}

	if len(requestBody.Data) == 0 {
		return c.Status(http.StatusBadRequest).SendString("No data provided")
	}

	err := repositories.InsertAlarm(requestBody.Data)
	if err != nil {
		log.Println("Error inserting alarm data:", err)
		return c.Status(http.StatusInternalServerError).SendString("Failed to insert alarms")
	}

	return c.Status(http.StatusCreated).JSON(requestBody.Data)
}

func (h handlers) GetAlarms(c *fiber.Ctx) error {
	var filter models.FilterAlarm

	err := c.QueryParser(&filter)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	alarms, err := repositories.GetAlarms(filter)
	if err != nil {
		log.Println("Error retrieving alarms:", err)
		return c.Status(http.StatusInternalServerError).SendString("Failed to retrieve alarms")
	}

	return c.JSON(fiber.Map{
		"data":  alarms,
		"count": len(alarms),
	})
}

func (h handlers) UpdateAlarmHandler(c *fiber.Ctx) error {
	alarmID := c.Params("alarm_id")

	var requestBody struct {
		Eventstatus string `json:"eventstatus"` // Change to eventstatus
		UpdatedBy   string `json:"updated_by"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(http.StatusBadRequest).SendString("Invalid request data")
	}

	alarm, err := repositories.GetAlarmByID(alarmID)
	if err != nil {
		log.Println("Error retrieving alarm:", err)
		return c.Status(http.StatusNotFound).SendString("Alarm not found")
	}

	if requestBody.Eventstatus != "" {
		alarm.Eventstatus = requestBody.Eventstatus // Change to eventstatus
	}

	if requestBody.UpdatedBy != "" {
		alarm.UpdatedBy = requestBody.UpdatedBy
	}

	alarm.UpdatedAt = time.Now()

	err = repositories.UpdateAlarm(alarm)
	if err != nil {
		log.Println("Error updating alarm:", err)
		return c.Status(http.StatusInternalServerError).SendString("Failed to update alarm")
	}

	return c.Status(http.StatusOK).JSON(alarm)
}
