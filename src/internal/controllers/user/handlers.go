package user

import (
	"github.com/graphql-go/graphql"
)

// [QUERY] user
func (t *UserControllerBuilder) QueryHello() *graphql.Field {
	return &graphql.Field{
		Type:        userType,
		Description: "Get the user phrase",
		Resolve:     t.queryUser,
	}
}

// [MUTATION] createUser
func (t *UserControllerBuilder) CreateUser() *graphql.Field {
	return &graphql.Field{
		Type:        userType,
		Description: "Update the user",

		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},

		Resolve: t.createUserMutation,
	}
}
