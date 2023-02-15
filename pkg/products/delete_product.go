package products

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api-docker/pkg/common/models"
	"go-fiber-api-docker/pkg/common/utils"
)

func (h handler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		utils.Logger.Error(result.Error.Error())
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&product)
	return c.SendStatus(fiber.StatusOK)
}
