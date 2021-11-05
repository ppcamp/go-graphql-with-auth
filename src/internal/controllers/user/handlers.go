package user

import (
	"github.com/graphql-go/graphql"
	usermodels "github.com/ppcamp/go-graphql-with-auth/internal/models/user"
	"github.com/ppcamp/go-graphql-with-auth/internal/services/jwt"
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

// [MUTATION] editUser
func (t *UserControllerBuilder) EditUser() *graphql.Field {
	return &graphql.Field{
		Type:        userType,
		Description: "Edit an user basing on its id",

		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},

			"nick": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return t.handler.Request(p, &usermodels.UserMutationPayload{}, NewEditUserController())
		},
	}
}

// [MUTATION] login
func (t *UserControllerBuilder) Login() *graphql.Field {
	return &graphql.Field{
		Type:        loginType,
		Description: "Does a login and generate an authenticated jwt",

		Args: graphql.FieldConfigArgument{
			"nick": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			_, err := t.handler.Request(p, &usermodels.UserMutationPayload{}, NewLoginUserController())
			if err == nil {
				s := jwt.Session{
					LoginSession: &jwt.LoginSession{
						UserId: 1,
					},
				}
				return jwt.BuildToken(p.Context, s)
			}
			return nil, err
		},
	}
}
