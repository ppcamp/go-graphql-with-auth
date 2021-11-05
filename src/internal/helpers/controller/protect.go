package controller

import (
	"github.com/graphql-go/graphql"
	"github.com/ppcamp/go-graphql-with-auth/internal/services/jwt"
)

func AuthorizedOnly(p graphql.ResolveParams) (interface{}, error) {
	_, err := jwt.GetSession(p.Context)
	if err != nil {
		return nil, err
	} else {
		return graphql.Field{}, err
	}
}
