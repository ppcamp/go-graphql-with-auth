package user

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/controller"
	usermodels "github.com/ppcamp/go-graphql-with-auth/internal/models/user"
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
