package user

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/controller"
	"github.com/ppcamp/go-graphql-with-auth/internal/models/usermodels"
)

type CreateUserController struct {
	controller.TransactionControllerImpl
}

func (c *CreateUserController) Execute(pl interface{}) (result controller.ResponseController) {
	user := pl.(*usermodels.UserMutationPayload)
	result = controller.NewResponseController()

	response, err := c.Transaction.CreateUser(user)
	result.SetError(err)
	result.SetResponse(response)

	return
}

func NewCreateUserController() controller.TransactionController {
	return &CreateUserController{}
}
