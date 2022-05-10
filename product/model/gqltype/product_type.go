package gqltype

import "github.com/graphql-go/graphql"

func ProductType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"desc": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
			"stock": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})
}
