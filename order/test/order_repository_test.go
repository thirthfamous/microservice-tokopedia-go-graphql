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

func TestCreateProductSuccess(t *testing.T) {
	db := app.NewDBTest()
	db.Exec("TRUNCATE `order`")
	orderRepository := repository.NewOrderRepository()
	order := domain.Order{
		ProfileId:   123,
		DateOrdered: time.Now(),
		ProductId:   123,
		Status:      1,
	}

	orderRepository.CreateOrder(db, &order)
	foundProduct := domain.Order{}
	db.Where("id = ?", order.Id).First(&foundProduct)

	assert.Equal(t, foundProduct.Id, order.Id)
}

func TestFindAllProductSuccess(t *testing.T) {
	db := app.NewDBTest()
	db.Exec("TRUNCATE `order`;")
	orderRepository := repository.NewOrderRepository()
	order := domain.Order{
		ProfileId:   123,
		DateOrdered: time.Now(),
		ProductId:   123,
		Status:      1,
	}
	db.Create(&order)

	result := orderRepository.FindOrderById(db, 1)
	assert.NotNil(t, result)
}

func TestUpdateOrderStatusToPaidSuccess(t *testing.T) {
	db := app.NewDBTest()
	db.Exec("TRUNCATE `order`;")
	orderRepository := repository.NewOrderRepository()
	order := domain.Order{
		ProfileId:   123,
		DateOrdered: time.Now(),
		ProductId:   123,
		Status:      1,
	}
	db.Create(&order)

	result := orderRepository.UpdateOrderStatusToPaid(db, 1)
	assert.NotNil(t, result)
}
