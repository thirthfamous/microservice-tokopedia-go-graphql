package impl

import (
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"
	"thirthfamous/tokopedia-clone-go-graphql/model/gqltype"
	"thirthfamous/tokopedia-clone-go-graphql/repository"
	"thirthfamous/tokopedia-clone-go-graphql/service"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *gorm.DB
}

func NewProductService(productRepository repository.ProductRepository, db *gorm.DB) service.GraphQLService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                db,
	}
}

func (service *ProductServiceImpl) MutationType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createProduct": &graphql.Field{
				Type: gqltype.ProductType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"desc": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"stock": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					admin := p.Context.Value("admin")
					if admin == nil {
						return nil, &gqlerrors.Error{
							Message: "Unauthenticated",
						}
					}
					product := domain.Product{
						Name:  p.Args["name"].(string),
						Desc:  p.Args["desc"].(string),
						Price: p.Args["price"].(int),
						Stock: p.Args["stock"].(int),
					}

					service.ProductRepository.CreateProduct(service.DB, &product)

					return product, nil
				},
			},
		},
	})
}

// QueryType implements service.GraphQLService
func (service *ProductServiceImpl) QueryType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"findAllProduct": &graphql.Field{
				Type: graphql.NewList(gqltype.ProductType()),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					profile_id := p.Context.Value("profile_id")
					if profile_id == nil {
						return nil, &gqlerrors.Error{
							Message: "Unauthenticated",
						}
					}
					products := service.ProductRepository.FindAllProduct(service.DB)

					return products, nil
				},
			},
		},
	})
}
