package test

import (
	"testing"
	"thirthfamous/tokopedia-clone-go-graphql/app"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"
	"thirthfamous/tokopedia-clone-go-graphql/repository"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnectionSuccess(t *testing.T) {
	db := app.NewDBTest()
	assert.NoError(t, db.Error)
}

func TestCreatePaymentSuccess(t *testing.T) {
	db := app.NewDBTest()
	db.Exec("TRUNCATE `payment`")
	paymentRepository := repository.NewPaymentRepository()
	payment := domain.Payment{
		OrderId:  1,
		PaidDate: time.Now(),
		Status:   1,
	}

	paymentRepository.CreatePayment(db, &payment)
	foundProduct := domain.Payment{}
	db.Where("id = ?", payment.Id).First(&foundProduct)

	assert.Equal(t, foundProduct.Id, payment.Id)
}

func TestFindAllProductSuccess(t *testing.T) {
	db := app.NewDBTest()
	db.Exec("TRUNCATE `payment`;")
	paymentRepository := repository.NewPaymentRepository()
	payment := domain.Payment{
		OrderId:  1,
		PaidDate: time.Now(),
		Status:   1,
	}
	db.Create(&payment)

	result := paymentRepository.CustomerPay(db, 1)
	assert.Equal(t, 2, result.Status) // paid
}
