package controller

import (
	"github.com/graphql-go/graphql"
	postgres "github.com/ppcamp/go-graphql-with-auth/internal/repository"
	"github.com/sirupsen/logrus"
)

type TransactionFunction func(tr postgres.Transaction, p graphql.ResolveParams) (interface{}, error)

func (h *Handler) Transaction(p graphql.ResolveParams, fn TransactionFunction) (interface{}, error) {
	tr, err := h.Storage.StartTransaction()
	if err != nil {
		return nil, err
	}

	// ensure db transaction is closed on panic
	defer func() {
		if r := recover(); r != nil {
			tr.Rollback()
			panic(r)
		}
	}()

	r, err := fn(tr, p)
	if err != nil {
		if err := tr.Rollback(); err != nil {
			logrus.WithError(err).Error("failed to rollback the commit")
			return nil, err
		}
	}

	if err := tr.Commit(); err != nil {
		logrus.WithError(err).Error("failed to commit")
		return nil, err
	}
	return r, err
}
