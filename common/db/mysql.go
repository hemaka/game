package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlRun() *gorm.DB {

	dsn := "root:123456@tcp(127.0.0.1:3306)/game?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
