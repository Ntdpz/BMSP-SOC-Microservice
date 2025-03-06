package handlers

import (
	"bmsp-backend-service/models"

	"github.com/gofiber/fiber/v2"
)

func (h handlers) GetListFileFromBuzzebee(c *fiber.Ctx) error {
	fileList := h.service.GetListFileFromBuzzebee()

	return c.JSON(fiber.Map{"data": fileList})
}

func (h handlers) GetListFileFromBuzzebeeStat(c *fiber.Ctx) error {
	lenFile, lenSentDb, lenWaitingDb := h.service.GetLenFileAndLenDBFromBuzzebee()

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"lenFile":      lenFile,
			"lenSentDb":    lenSentDb,
			"lenWaitingDb": lenWaitingDb,
		}},
	)
}

func (h handlers) GetDocumentListFromBuzzebee(c *fiber.Ctx) error {
	documents, err := h.service.GetDocumentListFromBuzzebee()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": documents})
}

func (h handlers) CreateDocumentBuzzebee(c *fiber.Ctx) error {

	var document models.Document

	if err := c.BodyParser(&document); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	document.Customer = "buzzebee"
	document.Status = "waiting"

	if err := h.service.CreateDocument(document); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": document})
}

func (h handlers) CreateDocumentBuzzebeeMultiple(c *fiber.Ctx) error {

	var document []models.Document

	if err := c.BodyParser(&document); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	for i := 0; i < len(document); i++ {
		document[i].Customer = "buzzebee"
		document[i].Status = "waiting"

		if err := h.service.CreateDocument(document[i]); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.JSON(fiber.Map{"data": document})
}
