package repository

import (
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(db *gorm.DB, order *domain.Order)
	FindOrderById(db *gorm.DB, orderId int) domain.Order
	FindOrderByProfileId(db *gorm.DB, orderId int) []domain.Order
}
