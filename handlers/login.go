package handlers

import (
	"bmsp-backend-service/models"

	"github.com/gofiber/fiber/v2"
)

func (h handlers) Login(c *fiber.Ctx) error {

	// get username and password from request body json
	request := new(models.RequestLogin)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := h.service.Login(request.Username, request.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	return c.JSON(fiber.Map{"token": token})

}
