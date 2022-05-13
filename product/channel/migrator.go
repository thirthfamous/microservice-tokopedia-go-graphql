package channel

import (
	"database/sql"
	"fmt"
	"thirthfamous/tokopedia-clone-go-graphql/helper"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"
	"thirthfamous/tokopedia-clone-go-graphql/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate() {

	/** CONNECT TO THE MYSQL */
	create_db, err := sql.Open("mysql", "root:123@tcp(db:3306)/")
	if err != nil {
		panic(err)
	}
	defer create_db.Close()

	/** CREATE THE DATABASES */
	_, err = create_db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", "tokopedia_product"))
	if err != nil {
		panic(err)
	}

	_, err = create_db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", "tokopedia_product_test"))
	if err != nil {
		panic(err)
	}

	utils.InitEnvironment()
	create_table := NewDB()
	create_table_test := NewDBTest()

	/** MIGRATE THE TABLES */
	create_table.AutoMigrate(
		domain.Product{},
	)
	create_table_test.AutoMigrate(
		domain.Product{},
	)
	fmt.Println("Migrate finished")
}

func NewDB() *gorm.DB {
	dsn := "root:123@tcp(db:3306)/tokopedia_product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	return db
}

func NewDBTest() *gorm.DB {
	dsn := "root:123@tcp(db:3306)/tokopedia_product_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	return db
}
