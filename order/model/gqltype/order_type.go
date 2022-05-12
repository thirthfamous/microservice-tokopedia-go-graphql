package gqltype

import "github.com/graphql-go/graphql"

func OrderType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"profile_id": &graphql.Field{
				Type: graphql.Int,
			},
			"product_id": &graphql.Field{
				Type: graphql.Int,
			},
			"date_ordered": &graphql.Field{
				Type: graphql.DateTime,
			},
			"status": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})
}
