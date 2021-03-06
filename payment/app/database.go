package app

import (
	"thirthfamous/tokopedia-clone-go-graphql/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "root:123@tcp(db:3306)/tokopedia_payment?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	return db
}

func NewDBTest() *gorm.DB {
	dsn := "root:123@tcp(127.0.0.1:3306)/tokopedia_payment_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	return db
}
