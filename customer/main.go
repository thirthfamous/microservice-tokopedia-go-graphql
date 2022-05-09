package main

import (
	"net/http"
	"thirthfamous/tokopedia-clone-go-graphql/app"
	"thirthfamous/tokopedia-clone-go-graphql/repository"
	userService "thirthfamous/tokopedia-clone-go-graphql/service/impl"
	"thirthfamous/tokopedia-clone-go-graphql/utils"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

func main() {
	utils.InitEnvironment()
	db := app.NewDB()

	userRepository := repository.NewUserRepository()
	userService := userService.NewUserService(userRepository, db)

	var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    userService.QueryType(),
		Mutation: userService.MutationType(),
	})
	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":3001", nil)

}
