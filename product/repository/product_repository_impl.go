package repository

import (
	"thirthfamous/tokopedia-clone-go-graphql/helper"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (*ProductRepositoryImpl) CreateProduct(db *gorm.DB, product *domain.Product) {
	result := db.Create(&product)
	helper.PanicIfError(result.Error)
}

func (*ProductRepositoryImpl) FindAllProduct(db *gorm.DB) []domain.Product {
	product := []domain.Product{}
	result := db.Find(&product)
	helper.PanicIfError(result.Error)
	return product
}
