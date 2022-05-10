package service

import "github.com/graphql-go/graphql"

type GraphQLService interface {
	QueryType() *graphql.Object
	MutationType() *graphql.Object
}
