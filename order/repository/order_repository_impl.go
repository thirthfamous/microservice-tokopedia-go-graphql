package repository

import (
	"fmt"
	"thirthfamous/tokopedia-clone-go-graphql/helper"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
}

// CreateProduct implements OrderRepository
func (*OrderRepositoryImpl) CreateOrder(db *gorm.DB, order *domain.Order) {
	result := db.Create(&order)
	helper.PanicIfError(result.Error)
}

func (*OrderRepositoryImpl) UpdateOrderStatusToPaid(db *gorm.DB, orderId int) domain.Order {
	result := db.Model(domain.Order{}).Where("id = ?", orderId).Update("status", 2) //paid
	helper.PanicIfError(result.Error)
	order := domain.Order{}
	db.Where("id = ?", orderId).Find(&order)
	return order
}

// FindProductById implements OrderRepository
func (*OrderRepositoryImpl) FindOrderById(db *gorm.DB, orderId int) domain.Order {
	order := domain.Order{}
	result := db.Where("id = ?", orderId).First(&order)
	fmt.Println(order)
	helper.PanicIfError(result.Error)
	return order
}

func (*OrderRepositoryImpl) FindOrderByProfileId(db *gorm.DB, profileId int) []domain.Order {
	order := []domain.Order{}
	fmt.Println(profileId)
	result := db.Where("profile_id = ?", profileId).Find(&order)
	fmt.Println(order)
	helper.PanicIfError(result.Error)
	return order
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}
