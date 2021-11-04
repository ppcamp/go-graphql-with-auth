package user

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/controller"
	"github.com/ppcamp/go-graphql-with-auth/internal/models/usermodels"
)

type QueryUserController struct {
	controller.TransactionControllerImpl
}

func (c *QueryUserController) Execute(pl interface{}) (result controller.ResponseController) {
	result = controller.NewResponseController()
	filter := pl.(*usermodels.UserQueryPayload)

	users, err := c.Transaction.FindUsers(filter)
	result.SetError(err)
	result.SetResponse(users)
	return
}

func NewQueryUserController() controller.TransactionController {
	return &QueryUserController{}
}
