package user

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/controller"
)

type QueryUserController struct {
	controller.BaseControllerImpl
}

func (c *QueryUserController) Execute(pl interface{}) (result controller.ResponseController) {
	result = controller.NewResponseController()
	return
}

func NewQueryUserController() controller.BaseController {
	return &QueryUserController{}
}
