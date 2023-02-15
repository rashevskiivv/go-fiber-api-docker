package products

import (
	"go-fiber-api-docker/pkg/common/models"
	"go-fiber-api-docker/pkg/common/utils"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		utils.Logger.Error(result.Error.Error())
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&product)
}
