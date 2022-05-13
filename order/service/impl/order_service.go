package impl

import (
	"fmt"
	"os"
	"thirthfamous/tokopedia-clone-go-graphql/messagebroker"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"
	"thirthfamous/tokopedia-clone-go-graphql/model/gqltype"
	"thirthfamous/tokopedia-clone-go-graphql/repository"
	"thirthfamous/tokopedia-clone-go-graphql/service"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	OrderRepository repository.OrderRepository
	DB              *gorm.DB
}

func NewProductService(productRepository repository.OrderRepository, db *gorm.DB) service.GraphQLService {
	return &ProductServiceImpl{
		OrderRepository: productRepository,
		DB:              db,
	}
}

func (service *ProductServiceImpl) MutationType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createOrder": &graphql.Field{
				Type: gqltype.OrderType(),
				Args: graphql.FieldConfigArgument{
					"profile_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"product_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					profile_id := p.Context.Value("profile_id")
					if profile_id == nil {
						return nil, &gqlerrors.Error{
							Message: "Unauthenticated",
						}
					}
					order := domain.Order{
						ProfileId:   p.Args["profile_id"].(int),
						ProductId:   p.Args["product_id"].(int),
						DateOrdered: time.Now(),
						Status:      1, // order created, need to pay
					}

					service.OrderRepository.CreateOrder(service.DB, &order)
					fmt.Println("Creating Order")
					fmt.Printf("os get env : %v", os.Getenv("TESTING"))
					fmt.Println(order)
					if os.Getenv("TESTING") == "false" {
						fmt.Println("Sending Create Payment Message")
						messagebroker.SendToMessageQueue("CreatePayment", order.Id)
					}

					return order, nil
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
			"findOrderById": &graphql.Field{
				Type: gqltype.OrderType(),
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
					order := service.OrderRepository.FindOrderById(service.DB, p.Args["id"].(int))
					return order, nil
				},
			},
			"findAllOrderByProfileId": &graphql.Field{
				Type: graphql.NewList(gqltype.OrderType()),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					profileId := p.Context.Value("profile_id")
					if profileId == nil {
						return nil, &gqlerrors.Error{
							Message: "Unauthenticated",
						}
					}
					product := service.OrderRepository.FindOrderByProfileId(service.DB, int(profileId.(float64)))
					return product, nil
				},
			},
		},
	},
	)
}
