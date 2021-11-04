package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ppcamp/go-graphql-with-auth/internal/models/usermodels"
)

// [QUERY] user
func (t *UserControllerBuilder) QueryUsers() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(userType),
		Description: "Get all users",

		Args: graphql.FieldConfigArgument{
			"nick": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"skip": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"take": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return t.handler.Request(p, &usermodels.UserQueryPayload{}, NewQueryUserController())
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
			return t.handler.Request(p, &usermodels.UserMutationPayload{}, NewCreateUserController())
		},
	}
}
