package user

import (
	"github.com/graphql-go/graphql"
)

func (t *UserControllerBuilder) createUserMutation(params graphql.ResolveParams) (interface{}, error) {
	name := params.Args["name"].(string)

	if name != "" {
		user.Name = name
	}

	return user, nil
}
