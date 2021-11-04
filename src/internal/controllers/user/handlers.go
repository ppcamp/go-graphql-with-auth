package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ppcamp/go-graphql-with-auth/internal/models"
)

// [QUERY] user
func (t *UserControllerBuilder) QueryHello() *graphql.Field {
	return &graphql.Field{
		Type:        userType,
		Description: "Get the user phrase",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return t.handler.Request(p, nil, NewQueryUserController())
		},
	}
}

// [MUTATION] createUser
func (t *UserControllerBuilder) CreateUser() *graphql.Field {
	return &graphql.Field{
		Type:        userType,
		Description: "Create a new user",

		Args: graphql.FieldConfigArgument{
			"nick": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return t.handler.Request(p, &models.User{}, NewCreateUserController())
		},
	}
}
