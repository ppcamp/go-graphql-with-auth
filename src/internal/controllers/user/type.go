package user

import (
	"time"

	"github.com/graphql-go/graphql"
)

type User struct {
	Id       *int64     `json:"id,omitempty"`
	Name     *string    `json:"name,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
}

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"birthday": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)