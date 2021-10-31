package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ppcamp/go-graphql-with-auth/internal/models"
	postgres "github.com/ppcamp/go-graphql-with-auth/internal/repository"
)

func (t *UserControllerBuilder) createUserMutation(
	tr postgres.Transaction,
	params graphql.ResolveParams,
) (interface{}, error) {
	user := models.User{
		Password: params.Args["password"].(string),
		Nick:     params.Args["nick"].(string),
		Email:    params.Args["email"].(string),
	}

	user, err := tr.CreateUser(&user)
	return user, err
}
