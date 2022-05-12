package impl

import (
	"fmt"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"
	"thirthfamous/tokopedia-clone-go-graphql/model/gqltype"
	"thirthfamous/tokopedia-clone-go-graphql/repository"
	"thirthfamous/tokopedia-clone-go-graphql/service"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"gorm.io/gorm"
)

type PaymentServiceImpl struct {
	PaymentRepository repository.PaymentRepository
	DB                *gorm.DB
}

func NewPaymentService(productRepository repository.PaymentRepository, db *gorm.DB) service.GraphQLService {
	return &PaymentServiceImpl{
		PaymentRepository: productRepository,
		DB:                db,
	}
}

func (service *PaymentServiceImpl) MutationType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createPayment": &graphql.Field{
				Type: gqltype.PaymentType(),
				Args: graphql.FieldConfigArgument{
					"order_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					order := domain.Payment{
						OrderId: p.Args["order_id"].(int),
						Status:  1,
					}

					service.PaymentRepository.CreatePayment(service.DB, &order)
					return order, nil
				},
			},
			"customerPay": &graphql.Field{
				Type: gqltype.PaymentType(),
				Args: graphql.FieldConfigArgument{
					"order_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					order := service.PaymentRepository.CustomerPay(service.DB, p.Args["order_id"].(int))
					fmt.Println(order)
					return order, nil
				},
			},
		},
	})
}

// QueryType implements service.GraphQLService
func (service *PaymentServiceImpl) QueryType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"findAllPayment": &graphql.Field{
				Type: gqltype.PaymentType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Context.Value("profile_id") == nil {
						return nil, &gqlerrors.Error{
							Message: "Unauthenticated",
						}
					}
					// order := service.PaymentRepository.FindAllPayment(service.DB, p.Args["id"].(int))
					return domain.Payment{}, nil
				},
			},
			"findAllOrderByProfileId": &graphql.Field{
				Type: graphql.NewList(gqltype.PaymentType()),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					profileId := p.Context.Value("profile_id")
					if profileId == nil {
						return nil, &gqlerrors.Error{
							Message: "Unauthenticated",
						}
					}
					// product := service.PaymentRepository.FindOrderByProfileId(service.DB, int(profileId.(float64)))
					return domain.Payment{}, nil
				},
			},
		},
	},
	)
}
