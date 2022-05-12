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
	productService "thirthfamous/tokopedia-clone-go-graphql/service/impl"
	"time"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupRouter() (http.Handler, *gorm.DB) {
	db := app.NewDBTest()
	db.Exec("TRUNCATE `order`;")
	productRepository := repository.NewOrderRepository()
	productService := productService.NewProductService(productRepository, db)

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

func TestCreateOrderEndpointSuccess(t *testing.T) {
	router, _ := setupRouter()
	profileId, _ := helper.GenerateToken(1)

	requestBody := strings.NewReader(`mutation
	RootMutation {
	  createOrder(profile_id:1, product_id:1){
		profile_id
		product_id
		date_ordered
		status
	  }
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3002/graphql", requestBody)
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

	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["createOrder"].(map[string]interface{})["profile_id"].(float64)))
}

func TestFindAllProductEndpointSuccess(t *testing.T) {
	router, db := setupRouter()

	order := domain.Order{
		ProfileId:   123,
		DateOrdered: time.Now(),
		ProductId:   123,
		Status:      1,
	}
	db.Create(&order)

	requestBody := strings.NewReader(`query
	RootQuery {
	  findOrderById(id: 1){
		id
		profile_id
		product_id
		date_ordered
		status
	  }
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/graphql", requestBody)
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

	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["findOrderById"].(map[string]interface{})["id"].(float64)))
}
