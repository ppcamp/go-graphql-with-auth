package controller

import (
	"github.com/graphql-go/graphql"
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/validators"
)

type BaseFunction func(pl interface{}) (interface{}, error)

func (h *Handler) Base(p graphql.ResolveParams, pl interface{}, fn BaseFunction) (interface{}, error) {
	err := validators.ShouldBind(p.Args, pl)
	if err != nil {
		return nil, err
	}

	return fn(pl)
}
