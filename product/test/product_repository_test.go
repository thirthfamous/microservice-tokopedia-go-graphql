package test

import (
	"testing"
	"thirthfamous/tokopedia-clone-go-graphql/app"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"
	"thirthfamous/tokopedia-clone-go-graphql/repository"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnectionSuccess(t *testing.T) {
	db := app.NewDBTest()
	assert.NoError(t, db.Error)
}

func TestCreateProductSuccess(t *testing.T) {
	db := app.NewDBTest()
	db.Exec("TRUNCATE product;")
	productRepository := repository.NewProductRepository()
	product := domain.Product{
		Name:  "PS4",
		Desc:  "Main console favorit semua orang",
		Price: 1000,
		Stock: 5,
	}

	productRepository.CreateProduct(db, &product)
	foundProduct := domain.Product{}
	db.Where("id = ?", product.Id).First(&foundProduct)

	assert.Equal(t, foundProduct.Id, product.Id)
}

func TestFindAllProductSuccess(t *testing.T) {
	db := app.NewDBTest()
	db.Exec("TRUNCATE product;")
	productRepository := repository.NewProductRepository()
	product := domain.Product{
		Name:  "PS4",
		Desc:  "Main console favorit semua orang",
		Price: 1000,
		Stock: 5,
	}
	db.Create(&product)

	result := productRepository.FindAllProduct(db)
	assert.NotNil(t, result)
}
