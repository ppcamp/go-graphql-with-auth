package user

import (
	"time"

	"github.com/graphql-go/graphql"
)

func (t *UserControllerBuilder) createUserMutation(params graphql.ResolveParams) (interface{}, error) {
	user := User{
		Name:     params.Args["name"].(*string),
		Birthday: params.Args["birthday"].(*time.Time),
	}
	return user, nil
}
