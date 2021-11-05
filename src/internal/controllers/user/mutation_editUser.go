package user

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/controller"
	usermodels "github.com/ppcamp/go-graphql-with-auth/internal/models/user"
)

type EditUserController struct {
	controller.TransactionControllerImpl
}

func (c *EditUserController) Execute(pl interface{}) (result controller.ResponseController) {
	user := pl.(*usermodels.UserMutationPayload)
	result = controller.NewResponseController()

	response, err := c.Transaction.CreateUser(user)
	result.SetError(err)
	result.SetResponse(response)

	return
}

func NewEditUserController() controller.TransactionController {
	return &EditUserController{}
}
