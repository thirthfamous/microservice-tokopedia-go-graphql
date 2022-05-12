package repository

import (
	"thirthfamous/tokopedia-clone-go-graphql/helper"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"

	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
}

func NewPaymentRepository() PaymentRepository {
	return &PaymentRepositoryImpl{}
}

// CreateProduct implements OrderRepository
func (*PaymentRepositoryImpl) CreatePayment(db *gorm.DB, order *domain.Payment) {
	result := db.Create(&order)
	helper.PanicIfError(result.Error)
}

// FindProductById implements OrderRepository
func (*PaymentRepositoryImpl) CustomerPay(db *gorm.DB, orderId int) domain.Payment {
	result := db.Model(domain.Payment{}).Where("id = ?", orderId).Update("status", 2) //paid
	helper.PanicIfError(result.Error)
	payment := domain.Payment{}
	db.Where("id = ?", orderId).Find(&payment)
	return payment
}
