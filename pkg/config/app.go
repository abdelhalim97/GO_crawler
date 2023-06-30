package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_crawler?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err) //we use it to fail fast on errors
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
