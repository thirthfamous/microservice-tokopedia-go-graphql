package channel

import (
	"database/sql"
	"fmt"
	"thirthfamous/tokopedia-clone-go-graphql/app"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"
	"thirthfamous/tokopedia-clone-go-graphql/utils"
)

func Migrate() {

	/** CONNECT TO THE MYSQL */
	create_db, err := sql.Open("mysql", "root:123@tcp(db:3306)/")
	if err != nil {
		panic(err)
	}
	defer create_db.Close()

	/** CREATE THE DATABASES */
	_, err = create_db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", "tokopedia_payment"))
	if err != nil {
		panic(err)
	}

	_, err = create_db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", "tokopedia_payment_test"))
	if err != nil {
		panic(err)
	}

	utils.InitEnvironment()
	create_table := app.NewDB()
	create_table_test := app.NewDBTest()

	/** MIGRATE THE TABLES */
	create_table.AutoMigrate(
		domain.Payment{},
	)
	create_table_test.AutoMigrate(
		domain.Payment{},
	)
	fmt.Println("Migrate finished")
}
