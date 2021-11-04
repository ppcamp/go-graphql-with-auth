package controller

import postgres "github.com/ppcamp/go-graphql-with-auth/internal/repository"

// TransactionController implements a controller with database support.
// This controller makes all request using a sql transaction. Therefore,
// if some error occurrs (or if the status code is greater or equal than 400)
// it'll rollback all queries maded.
//
// See [api/controller/response.go] and [api/helpers/controller.go]
// See also BaseControllerImpl
//
// Example:
// 	func NewBasicGet() controller.TransactionController() {
//		return &basicGet{}
//	}
type TransactionController interface {
	BaseController
	SetTransaction(tr postgres.Transaction)
}

// TransactionControllerImpl should be used inside a typing controller
// definition.
// This will ensures that you can return a TransactionController object.
// See TransactionController
//
// Example:
//	type basicGet struct {
//		controller.TransactionControllerImpl
//	}
type TransactionControllerImpl struct {
	BaseControllerImpl

	Transaction postgres.Transaction
}

func (u *TransactionControllerImpl) SetTransaction(tr postgres.Transaction) {
	u.Transaction = tr
}
