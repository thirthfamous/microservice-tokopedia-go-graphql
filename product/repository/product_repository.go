package repository

import (
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(db *gorm.DB, product *domain.Product)
	FindAllProduct(db *gorm.DB) []domain.Product
}
