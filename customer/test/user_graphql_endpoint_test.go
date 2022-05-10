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
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"
	"thirthfamous/tokopedia-clone-go-graphql/repository"
	userService "thirthfamous/tokopedia-clone-go-graphql/service/impl"

	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/graphql-go/graphql"
)

func setupRouter() (http.Handler, *gorm.DB) {
	db := app.NewDBTest()
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE customer_user;")
	db.Exec("TRUNCATE customer_profile")
	userRepository := repository.NewUserRepository()
	userService := userService.NewUserService(userRepository, db)

	var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    userService.QueryType(),
		Mutation: userService.MutationType(),
	})

	http.DefaultServeMux = new(http.ServeMux)

	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	return h, db
}

func TestSignupSuccess(t *testing.T) {
	router, _ := setupRouter()

	requestBody := strings.NewReader(`mutation
	RootMutation {
	  createUser(username: "Farhassn", password: "123", name: "Andrew", address: "Sukabumi"){
		token
		profile {
		  id
		  name
		  address
		}
	  }
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/graphql", requestBody)
	request.Header.Add("Content-Type", "application/graphql")
	// request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "Andrew", responseBody["data"].(map[string]interface{})["createUser"].(map[string]interface{})["profile"].(map[string]interface{})["name"])
}

func TestSignupDuplicateEntryFailed(t *testing.T) {
	router, db := setupRouter()
	user := domain.User{
		Username: "Farhassn",
		Password: "123",
		Profile: domain.Profile{
			Name:    "Andrew",
			Address: "Sukabumi",
		},
	}
	db.Create(&user)

	requestBody := strings.NewReader(`mutation
	RootMutation {
	  createUser(username: "Farhassn", password: "123", name: "Andrew", address: "Sukabumi"){
		token
		profile {
		  id
		  name
		  address
		}
	  }
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/graphql", requestBody)
	request.Header.Add("Content-Type", "application/graphql")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.NotNil(t, responseBody["errors"])
}

func TestLoginSuccess(t *testing.T) {
	router, db := setupRouter()
	password, _ := helper.HashPassword("123")
	user := domain.User{
		Username: "Farhassn",
		Password: password,
		Profile: domain.Profile{
			Name:    "Andrew",
			Address: "Sukabumi",
		},
	}
	db.Create(&user)

	requestBody := strings.NewReader(`mutation
	RootMutation {
	  login(username: "Farhassn", password: "123")
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/graphql", requestBody)
	request.Header.Add("Content-Type", "application/graphql")
	// request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.NotNil(t, responseBody["data"].(map[string]interface{})["login"])
}
