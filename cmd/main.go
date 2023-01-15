package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api-docker/pkg/common/config"
	"go-fiber-api-docker/pkg/common/db"
	"go-fiber-api-docker/pkg/products"
	"log"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()
	products.RegisterRoutes(app, h)
	err = app.Listen(c.Port)
	if err != nil {
		log.Fatalln("Failed at listen", err)
	}
}