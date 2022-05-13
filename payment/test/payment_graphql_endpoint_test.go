package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"thirthfamous/tokopedia-clone-go-graphql/app"
	"thirthfamous/tokopedia-clone-go-graphql/helper"
	"thirthfamous/tokopedia-clone-go-graphql/middleware"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"
	"thirthfamous/tokopedia-clone-go-graphql/repository"
	paymentService "thirthfamous/tokopedia-clone-go-graphql/service/impl"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupRouter() (http.Handler, *gorm.DB) {
	db := app.NewDBTest()
	db.Exec("TRUNCATE `payment`;")
	productRepository := repository.NewPaymentRepository()
	productService := paymentService.NewPaymentService(productRepository, db)

	var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Mutation: productService.MutationType(),
		Query:    productService.QueryType(),
	})

	http.DefaultServeMux = new(http.ServeMux)

	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	return h, db
}

func TestCreatePaymentEndpointSuccess(t *testing.T) {
	router, _ := setupRouter()
	profileId, _ := helper.GenerateToken(1)

	requestBody := strings.NewReader(`mutation
	RootMutation {
	  createPayment(order_id:1){
		id
		order_id
		paid_date
		status
	  }
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/payment/graphql", requestBody)
	request.Header.Add("Content-Type", "application/graphql")
	request.Header.Add("Authorization", "Bearer "+profileId)

	recorder := httptest.NewRecorder()

	middleware.EnforceJSONHandler(middleware.AuthMiddleware(router)).ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["createPayment"].(map[string]interface{})["order_id"].(float64)))
}

func TestPayCustomerEndpointSuccess(t *testing.T) {
	router, db := setupRouter()

	payment := domain.Payment{
		OrderId: 1,
		Status:  1,
	}
	db.Create(&payment)

	requestBody := strings.NewReader(`mutation
	RootMutation {
		customerPay(order_id:1){
		id
		order_id
		paid_date
		status
	  }
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/payment/graphql", requestBody)
	request.Header.Add("Content-Type", "application/graphql")
	profileId, _ := helper.GenerateToken(1)
	request.Header.Add("Authorization", "Bearer "+profileId)

	recorder := httptest.NewRecorder()

	middleware.EnforceJSONHandler(middleware.AuthMiddleware(router)).ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["customerPay"].(map[string]interface{})["order_id"].(float64)))
}
