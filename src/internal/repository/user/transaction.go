package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/ppcamp/go-graphql-with-auth/internal/models"
)

type UserTransaction struct {
	tx *sqlx.Tx
}

type UserStorage interface {
	CreateUser(payload *models.User) (user models.User, err error)
}

func NewTransaction(tx *sqlx.Tx) *UserTransaction {
	return &UserTransaction{
		tx: tx,
	}
}
