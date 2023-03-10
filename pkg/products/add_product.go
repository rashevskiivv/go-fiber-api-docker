package products

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api-docker/pkg/common/models"
	"go-fiber-api-docker/pkg/common/utils"
)

type AddProductRequestBody struct {
	Name  string `json:"name"`
	Stock int32  `json:"stock"`
	Price int32  `json:"price"`
}

func (h handler) AddProduct(c *fiber.Ctx) error {
	body := AddProductRequestBody{}

	// parse body, attach to AddProductRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		utils.Logger.Error(err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var product models.Product

	product.Name = body.Name
	product.Stock = body.Stock
	product.Price = body.Price

	// insert new db entry
	if result := h.DB.Create(&product); result.Error != nil {
		utils.Logger.Error(result.Error.Error())
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&product)
}
