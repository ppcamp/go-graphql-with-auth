package user

import (
	"time"

	"github.com/graphql-go/graphql"
)

var user = User{
	Id:       0,
	Name:     "SomeUser",
	Birthday: time.Now(),
}

func (t *UserControllerBuilder) queryUser(p graphql.ResolveParams) (interface{}, error) {
	return user, nil
}
