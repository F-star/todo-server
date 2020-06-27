package main

import (
	"todo/config"
	initapp "todo/init"
	"todo/models"
	"todo/redis"
	"todo/routes"

	"github.com/jinzhu/gorm"
)

func main() {
	initapp.SetEnv()
	var err error
	config.DB, err = gorm.Open(
		"mysql",
		config.DbURL(config.BuildDBConfig()),
	)
	if err != nil {
		panic(err)
	}

	config.DB.LogMode(true) // open gorm log
	redis.StartRedis()

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Todo{})

	r := routes.SetupRouter()
	r.Run()
}
