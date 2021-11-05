package user

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/controller"
	loginmodels "github.com/ppcamp/go-graphql-with-auth/internal/models/login"
	usermodels "github.com/ppcamp/go-graphql-with-auth/internal/models/user"
)

type LoginUserController struct {
	controller.TransactionControllerImpl
}

func (c *LoginUserController) Execute(pl interface{}) (result controller.ResponseController) {
	user := pl.(*usermodels.UserMutationPayload)
	result = controller.NewResponseController()

	_, err := c.Transaction.FindUserWithPassword(user)

	r := loginmodels.TokenResponse{
		Token:   "asfasf",
		Expires: "asfas",
	}
	result.SetError(err)
	result.SetResponse(r)

	return
}

func NewLoginUserController() controller.TransactionController {
	return &LoginUserController{}
}
