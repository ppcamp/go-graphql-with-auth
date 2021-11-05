package user

import (
	"github.com/jmoiron/sqlx"
	usermodels "github.com/ppcamp/go-graphql-with-auth/internal/models/user"
)

type UserTransaction struct {
	tx *sqlx.Tx
}

type UserStorage interface {
	CreateUser(payload *usermodels.UserMutationPayload) (user usermodels.UserEntity, err error)
	EditUser(payload *usermodels.UserMutationPayload) (user usermodels.UserEntity, err error)
	FindUsers(filter *usermodels.UserQueryPayload) (users []usermodels.UserEntity, err error)
	FindUserWithPassword(filter *usermodels.UserMutationPayload) (user usermodels.UserEntity, err error)
}

func NewTransaction(tx *sqlx.Tx) *UserTransaction {
	return &UserTransaction{
		tx: tx,
	}
}
