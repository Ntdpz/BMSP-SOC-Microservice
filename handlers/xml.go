package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h handlers) CreateXML(c *fiber.Ctx) error {

	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if idInt <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id must be greater than 0"})
	}

	// get doc
	doc, err := h.service.GetDocumentFromBuzzebee(idInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if doc.Id == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error create document"})
	}

	// create xml
	if err := h.service.CreateXML(doc); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// update doc
	if err := h.service.UpdateDocument(doc.Id, "sent"); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": doc})
}
