package app

import (
	"github.com/graphql-go/graphql"
)

// [QUERY] app
func (t *AppController) QueryAppStatus() *graphql.Field {
	return &graphql.Field{
		Type:        appStatusType,
		Description: "Get the app status",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return t.handler.Transaction(p, t.queryPing)
		},
	}
}
