package handlers

import (
	"bmsp-backend-service/models"
	"bmsp-backend-service/repositories"
	"log"

	"github.com/gofiber/fiber/v2"
)

// CreateAlienvaultCaseHandler รับคำขอ POST และบันทึกข้อมูลลงในฐานข้อมูล
func CreateAlienvaultCaseHandler(c *fiber.Ctx) error {
	var requestBody struct {
		Data []models.AlienvaultCase `json:"data"`
	}

	// แปลงข้อมูลจาก body ของคำขอเป็น struct
	if err := c.BodyParser(&requestBody); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(400).SendString("Invalid request data")
	}

	// ตรวจสอบว่ามีข้อมูลใน array หรือไม่
	if len(requestBody.Data) == 0 {
		return c.Status(400).SendString("No data provided")
	}

	// บันทึกข้อมูลลงฐานข้อมูล
	err := repositories.InsertAlienvaultCase(requestBody.Data)
	if err != nil {
		return c.Status(500).SendString("Failed to insert cases")
	}

	// ส่งกลับข้อมูลที่ถูกบันทึก
	return c.Status(201).JSON(requestBody.Data)
}
