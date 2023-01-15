package db

import (
	"fmt"
	"go-fiber-api-docker/pkg/common/config"
	"go-fiber-api-docker/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(c *config.Config) *gorm.DB {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
