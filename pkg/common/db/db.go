package db

import (
	"fmt"
	"go-fiber-api-docker/pkg/common/config"
	"go-fiber-api-docker/pkg/common/models"
	"go-fiber-api-docker/pkg/common/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(c *config.Config) (db *gorm.DB, err error) {
	defer func() {
		if err == nil {
			utils.Logger.Info("DB initiated successfully")
		}
	}()

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return
	}
	err = db.AutoMigrate(&models.Product{})
	return
}
