package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/ppcamp/go-graphql-with-auth/internal/repository/status"
	"github.com/ppcamp/go-graphql-with-auth/internal/repository/user"
)

// Holds all objects that connects into postgresql
type transaction struct {
	*sqlx.Tx

	*status.StatusTransaction
	*user.UserTransaction
}

type Transaction interface {
	Commit() error
	Rollback() error

	user.UserStorage
	status.StatusStorage
}

// Starts a new transaction basing on a given storage transaction
func NewTransaction(tx *sqlx.Tx) Transaction {
	return &transaction{
		Tx: tx,

		StatusTransaction: status.NewTransaction(tx),
		UserTransaction:   user.NewTransaction(tx),
	}
}
