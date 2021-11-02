package user

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/models"
	postgres "github.com/ppcamp/go-graphql-with-auth/internal/repository"
)

func (t *UserControllerBuilder) createUserMutation(
	tr postgres.Transaction,
	pl interface{},
) (interface{}, error) {
	user := pl.(models.User)
	user, err := tr.CreateUser(&user)
	return user, err
}
