package db

import (
	"echo-gorm/config"
	"echo-gorm/model"

	// "gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	c := config.GetConfig()
	// dsn := fmt.Sprintf(
	// 	"%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	c.DB_USERNAME, c.DB_PASSWORD, c.DB_NAME)
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err = gorm.Open(sqlite.Open(c.DB_FILENAME), &gorm.Config{})

	// defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}

	db.AutoMigrate(&model.User{})
}

func DbManager() *gorm.DB {
	return db
}
