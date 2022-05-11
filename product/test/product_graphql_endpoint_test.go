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

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupRouter() (http.Handler, *gorm.DB) {
	db := app.NewDBTest()
	db.Exec("TRUNCATE product;")
	productRepository := repository.NewProductRepository()
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

func TestCreateProductEndpointSuccess(t *testing.T) {
	router, _ := setupRouter()

	requestBody := strings.NewReader(`mutation
	RootMutation {
	  createProduct(name: "PS4", desc: "Console kesukaan semua orang", price: 1000, stock: 5){
		name
		desc  
		price 
		stock 
	  }
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3002/graphql", requestBody)
	request.Header.Add("Content-Type", "application/graphql")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	middleware.EnforceJSONHandler(middleware.AuthMiddleware(router)).ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, "PS4", responseBody["data"].(map[string]interface{})["createProduct"].(map[string]interface{})["name"])
}

func TestFindAllProductEndpointSuccess(t *testing.T) {
	router, db := setupRouter()

	product := domain.Product{
		Name:  "PS4",
		Desc:  "Main console favorit semua orang",
		Price: 1000,
		Stock: 5,
	}
	db.Create(&product)

	requestBody := strings.NewReader(`query
	RootQuery {
	  findAllProduct{
		name
		desc  
		price 
		stock 
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

	assert.Equal(t, "PS4", responseBody["data"].(map[string]interface{})["findAllProduct"].([]interface{})[0].(map[string]interface{})["name"])
}
