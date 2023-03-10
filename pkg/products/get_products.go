package products

import (
	"go-fiber-api-docker/pkg/common/models"
	"go-fiber-api-docker/pkg/common/utils"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	if result := h.DB.Find(&products); result.Error != nil {
		utils.Logger.Error(result.Error.Error())
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&products)
}
