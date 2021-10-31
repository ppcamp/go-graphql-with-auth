package app

import (
	"github.com/graphql-go/graphql"
)

type AppStatus struct {
	Postgresql      bool  `json:"postgres,omitempty"`
	ConnectionDelay int64 `json:"delay,omitempty"`
}

var appStatusType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "App",
		Fields: graphql.Fields{
			"postgres": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "If the database is working or not",
			},
			"delay": &graphql.Field{
				Type:        graphql.Int,
				Description: "The amount of time that took to get the postgresql query",
			},
		},
	},
)
