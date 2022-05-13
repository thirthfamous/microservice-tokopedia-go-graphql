package main

import (
	"net/http"
	"thirthfamous/tokopedia-clone-go-graphql/app"
	"thirthfamous/tokopedia-clone-go-graphql/channel"
	"thirthfamous/tokopedia-clone-go-graphql/middleware"
	"thirthfamous/tokopedia-clone-go-graphql/repository"
	productService "thirthfamous/tokopedia-clone-go-graphql/service/impl"
	"thirthfamous/tokopedia-clone-go-graphql/utils"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

func main() {
	channel.Migrate()
	go channel.Receive()
	utils.InitEnvironment()
	db := app.NewDB()

	productRepository := repository.NewPaymentRepository()
	productService := productService.NewPaymentService(productRepository, db)

	var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    productService.QueryType(),
		Mutation: productService.MutationType(),
	})
	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	http.Handle("/graphql", middleware.EnforceJSONHandler(middleware.AuthMiddleware(h)))
	http.ListenAndServe(":3004", nil)

}
