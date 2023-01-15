package products

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api-docker/pkg/common/models"
)

type UpdateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int32  `json:"stock"`
	Price int32  `json:"price"`
}

func (h handler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateProductRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	product.Name = body.Name
	product.Stock = body.Stock
	product.Price = body.Price

	h.DB.Save(&product)
	return c.Status(fiber.StatusOK).JSON(&product)
}
