package impl

import (
	"thirthfamous/tokopedia-clone-go-graphql/helper"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"
	"thirthfamous/tokopedia-clone-go-graphql/model/gqltype"
	"thirthfamous/tokopedia-clone-go-graphql/repository"
	"thirthfamous/tokopedia-clone-go-graphql/service"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
}

func NewUserService(userRepository repository.UserRepository, db *gorm.DB) service.GraphQLService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *UserServiceImpl) MutationType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: gqltype.AuthPayloadType(),
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					password, _ := helper.HashPassword(p.Args["password"].(string))
					user := domain.User{
						Username: p.Args["username"].(string),
						Password: password,
						Profile: domain.Profile{
							Name:    p.Args["name"].(string),
							Address: p.Args["address"].(string),
						},
					}
					service.UserRepository.CreateUser(service.DB, &user)
					profile, _ := helper.GenerateToken(user.Profile)
					return domain.AuthPayload{
						Token:   profile,
						Profile: user.Profile,
					}, nil
				},
			},
			"login": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					user := domain.User{}
					service.UserRepository.FindByUsername(service.DB, p.Args["username"].(string), &user)
					if user == (domain.User{}) {
						return nil, &gqlerrors.Error{
							Message: "Akun tidak ditemukan",
						}
					}
					validPassword := helper.CheckPasswordHash(p.Args["password"].(string), user.Password)
					if !validPassword {
						return nil, &gqlerrors.Error{
							Message: "Username atau Password Salah",
						}
					}
					token, _ := helper.GenerateToken(user.Profile)
					return token, nil
				},
			},
		},
	},
	)
}

func (service *UserServiceImpl) QueryType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type: graphql.NewList(service.UserType()),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					user := []domain.User{}
					result := service.DB.Find(&user)
					if result.Error != nil {
						panic(result.Error)
					}
					return user, result.Error
				},
			},
			"user": &graphql.Field{
				Type: service.UserType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					user := domain.User{}
					result := service.DB.First(&user, p.Args["id"])
					if result.Error != nil {
						panic(result.Error)
					}
					return user, result.Error
				},
			},
		},
	},
	)
}
