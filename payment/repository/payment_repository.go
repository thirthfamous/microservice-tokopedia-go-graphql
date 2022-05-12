package repository

import (
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreatePayment(db *gorm.DB, payment *domain.Payment)
	CustomerPay(db *gorm.DB, paymentId int) domain.Payment
}
