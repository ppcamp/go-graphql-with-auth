package controller

import (
	"github.com/graphql-go/graphql"
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/validators"
	postgres "github.com/ppcamp/go-graphql-with-auth/internal/repository"
	"github.com/sirupsen/logrus"
)

type TransactionFunction func(tr postgres.Transaction, pl interface{}) (interface{}, error)

func (h *Handler) Transaction(p graphql.ResolveParams, pl interface{}, fn TransactionFunction) (interface{}, error) {
	if pl != nil {
		err := validators.ShouldBind(p.Args, pl)
		if err != nil {
			return nil, err
		}
	}

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

	r, err := fn(tr, pl)
	if err != nil {
		if err := tr.Rollback(); err != nil {
			logrus.WithError(err).Error("failed to rollback the commit")
		}
		return nil, err
	} else {
		if err := tr.Commit(); err != nil {
			logrus.WithError(err).Error("failed to commit")
		}
		return r, err
	}
}
