package user

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/controller"
	usermodels "github.com/ppcamp/go-graphql-with-auth/internal/models/user"
)

type LoginUserController struct {
	controller.TransactionControllerImpl
}

func (c *LoginUserController) Execute(pl interface{}) (result controller.ResponseController) {
	user := pl.(*usermodels.UserMutationPayload)
	result = controller.NewResponseController()

	_, err := c.Transaction.FindUserWithPassword(user)
	result.SetError(err)
	return
}

func NewLoginUserController() controller.TransactionController {
	return &LoginUserController{}
}
