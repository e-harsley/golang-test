package database

import (
	"context"
	"fmt"

	"github.com/go-play/config"
	"github.com/go-play/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Ctx context.Context

func Init() *gorm.DB {
	dsn := config.DB_USERNAME + ":" + config.DB_PASSWORD + "@tcp" + "(" + config.DB_HOST + ":" + config.DB_PORT + ")/" + "?" + "parseTime=true&loc=Local"
	fmt.Println("hello", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utils.PanicException(err, "database connection error!")
	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.DB_NAME + ";")

	dsn = config.DB_USERNAME + ":" + config.DB_PASSWORD + "@tcp" + "(" + config.DB_HOST + ":" + config.DB_PORT + ")/" + config.DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utils.PanicException(err, "database connection error!")

	migrate(db)

	return db
}
