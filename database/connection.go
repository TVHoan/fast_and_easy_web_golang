package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:1@tcp(127.0.0.1:3306)/go_web?charset=utf8mb4&parseTime=True&loc=Local"

func Db() (db *gorm.DB) {
	dbContext, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect database")
	}
	return dbContext

}
