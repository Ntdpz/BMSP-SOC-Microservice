package handlers

import "github.com/gofiber/fiber/v2"

func (h handlers) RootHandler(c *fiber.Ctx) error {
	return c.SendString("BMSP-SOC-Microservice started!")
}
