package status

import (
	"github.com/jmoiron/sqlx"
)

type StatusTransaction struct {
	tx *sqlx.Tx
}

type StatusStorage interface {
	Ping() (err error)
}

func NewTransaction(tx *sqlx.Tx) *StatusTransaction {
	return &StatusTransaction{
		tx: tx,
	}
}
