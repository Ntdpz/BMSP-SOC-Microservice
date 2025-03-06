package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h handlers) GetListFileFromHorizon(c *fiber.Ctx) error {
	fileList := h.service.GetListFileFromHorizon()

	return c.JSON(fiber.Map{"data": fileList})
}
