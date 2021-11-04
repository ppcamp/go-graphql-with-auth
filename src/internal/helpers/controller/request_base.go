package controller

import (
	"context"

	"github.com/ppcamp/go-graphql-with-auth/internal/services/jwt"
)

// BaseController should be used as returning type for a controller object.
// See BaseControllerImpl
//
// Example:
// 	func NewBasicGet() controller.BaseController() {
//		return &basicGet{}
//	}
type BaseController interface {
	SetSession(session jwt.Session)
	SetContext(ctx context.Context)

	// Execute is the function that will be executed by the controller
	//
	// Example
	//	type getAppStatus struct {
	//		controller.TransactionControllerImpl
	//	}
	//	func (m *getAppStatus) Execute(pl interface{}) (result controller.ResponseController) {}
	Execute(payload interface{}) ResponseController
}

// BaseControllerImpl should be used inside a typing controller definition.
// This will ensures that you can return a BaseController object.
// See BaseController
//
// Example:
//	type basicGet struct {
//		controller.BaseControllerImpl
//	}
type BaseControllerImpl struct {
	Session jwt.Session
	context context.Context
}

func (bc *BaseControllerImpl) SetContext(ctx context.Context) {
	bc.context = ctx
}

func (bc *BaseControllerImpl) SetSession(session jwt.Session) {
	bc.Session = session
}
