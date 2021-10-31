package user

import (
	"github.com/graphql-go/graphql"
)

var Phrase = "world"

// [QUERY] user
func (t *UserControllerBuilder) GetHello() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "Get the user phrase",
		Resolve:     t.getHello,
	}
}

// [MUTATION] createUser
func (t *UserControllerBuilder) CreateUser() *graphql.Field {
	return &graphql.Field{
		Type:        userType,
		Description: "Create new user",

		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"birthday": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
		},

		Resolve: t.createUserMutation,
	}
}
