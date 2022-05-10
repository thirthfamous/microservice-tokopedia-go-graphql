package main

import (
	"database/sql"
	"fmt"
	"thirthfamous/tokopedia-clone-go-graphql/app"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"
	"thirthfamous/tokopedia-clone-go-graphql/utils"
)

func main() {

	/** CONNECT TO THE MYSQL */
	create_db, err := sql.Open("mysql", "root@tcp(localhost:3306)/")
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
	create_table := app.NewDB()
	create_table_test := app.NewDBTest()

	/** MIGRATE THE TABLES */
	create_table.AutoMigrate(
		domain.Product{},
	)
	create_table_test.AutoMigrate(
		domain.Product{},
	)
	fmt.Println("Migrate finished")
}