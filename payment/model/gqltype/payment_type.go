package gqltype

import "github.com/graphql-go/graphql"

func PaymentType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"order_id": &graphql.Field{
				Type: graphql.Int,
			},
			"paid_date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"status": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})
}
