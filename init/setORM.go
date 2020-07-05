// 初始化数据库表格。
package initapp

import (
	"todo/config"
	"todo/models"

	"github.com/jinzhu/gorm"
)

func SetORM() {
	var err error
	config.DB, err = gorm.Open(
		"mysql",
		config.DbURL(config.BuildDBConfig()),
	)
	if err != nil {
		panic(err)
	}

	config.DB.LogMode(true) // open gorm log

	config.DB.AutoMigrate(&models.Todo{})
	config.DB.AutoMigrate(&models.User{})

}
