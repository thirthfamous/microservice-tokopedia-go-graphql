package gqltype

import "github.com/graphql-go/graphql"

func ProfileType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Profile",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func UserType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"profile": &graphql.Field{
				Type: ProfileType(),
			},
		},
	})
}

func AuthPayloadType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "AuthPayload",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"profile": &graphql.Field{
				Type: ProfileType(),
			},
		},
	})
}
