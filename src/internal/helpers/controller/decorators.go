package controller

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/validators"
	"github.com/ppcamp/go-graphql-with-auth/internal/services/jwt"
)

// Request will check the type of the baseController object
// according with the type will use a BaseController or a TransactionController
func (h *Handler) Request(
	p graphql.ResolveParams,
	payload interface{},
	ctrl BaseController,
) (interface{}, error) {
	if payload != nil {
		err := validators.ShouldBind(p.Args, payload)
		if err != nil {
			return nil, err
		}
	}

	ctrl.SetContext(p.Context)
	session, _ := jwt.GetSession(p.Context)
	ctrl.SetSession(session)

	switch ctrl := ctrl.(type) {
	case TransactionController:
		return h.transactionController(payload, ctrl)
	default:
		return h.baseController(payload, ctrl)
	}
}

// baseController is a simple controller. It doesn't implement the database and
// transactions setup
func (h *Handler) baseController(
	payload interface{},
	ctrl BaseController,
) (interface{}, error) {
	response := ctrl.Execute(payload)

	if err := response.GetError(); err != nil {
		return nil, err
	} else {
		return response.GetResponse(), nil
	}
}

// transactionController implements a middleware decorator that can be used to
// assign a status code and automatically parse the object
func (h *Handler) transactionController(
	payload interface{},
	ctrl TransactionController,
) (interface{}, error) {
	tr, err := h.Storage.StartTransaction()
	if err != nil {
		return nil, err
	}
	ctrl.SetTransaction(tr)

	// ensure db transaction is closed on panic
	defer func() {
		if r := recover(); r != nil {
			tr.Rollback()
			panic(r)
		}
	}()

	response := ctrl.Execute(payload)

	if err := response.GetError(); err != nil {
		if err := tr.Rollback(); err != nil {
			return nil, fmt.Errorf("%v : %v", response.GetError(), err)
		} else {
			return nil, response.GetError()
		}
	}

	if err := tr.Commit(); err != nil {
		return nil, err
	}

	return response.GetResponse(), err
}
