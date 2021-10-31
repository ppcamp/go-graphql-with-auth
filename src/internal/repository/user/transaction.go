package user

import (
	"github.com/jmoiron/sqlx"
)

type UserTransaction struct {
	tx *sqlx.Tx
}

type UserStorage interface {
}

func NewTransaction(tx *sqlx.Tx) *UserTransaction {
	return &UserTransaction{
		tx: tx,
	}
}
