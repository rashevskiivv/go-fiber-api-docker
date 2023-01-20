package products

import "github.com/gofiber/fiber/v2"

func (h handler) Welcome(c *fiber.Ctx) error {
	welcome := "Hi! Here's nothing interesting instead of /products. By the way we recommend you to check out our README file on https://github.com/rashevskiivv/go-fiber-api-docker"
	return c.Status(fiber.StatusOK).SendString(welcome)
}
