package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api-docker/pkg/common/config"
	"go-fiber-api-docker/pkg/common/db"
	"go-fiber-api-docker/pkg/common/utils"
	"go-fiber-api-docker/pkg/products"
)

func main() {
	utils.InitLogger()
	utils.Logger.Info("Server started")

	c, err := config.LoadConfig()
	if err != nil {
		utils.Logger.Error(err.Error())
	}

	h, err := db.Init(&c)
	if err != nil {
		utils.Logger.Error(err.Error())
	}

	app := fiber.New()
	products.RegisterRoutes(app, h)
	err = app.Listen(c.Port)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
}
