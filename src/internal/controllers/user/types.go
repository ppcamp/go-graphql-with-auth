package user

import (
	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "User's id",
			},
			"nick": &graphql.Field{
				Type:        graphql.String,
				Description: "User's nickname",
			},
			"email": &graphql.Field{
				Type:        graphql.String,
				Description: "User's email",
			},
			"updated_at": &graphql.Field{
				Type:        graphql.String,
				Description: "When the field was updated",
			},
		},
	},
)

var loginType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Login",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type:        graphql.String,
				Description: "The JWT token needed for use 'me'",
			},
			"expires": &graphql.Field{
				Type:        graphql.String,
				Description: "When the token will expire",
			},
		},
	},
)
