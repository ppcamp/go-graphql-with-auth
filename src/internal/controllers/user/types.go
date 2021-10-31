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
