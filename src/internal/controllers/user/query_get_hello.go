package user

import "github.com/graphql-go/graphql"

func (t *UserControllerBuilder) getHello(p graphql.ResolveParams) (interface{}, error) {
	return Phrase, nil
}
