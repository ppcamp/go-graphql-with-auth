package user

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/controller"
	"github.com/ppcamp/go-graphql-with-auth/internal/models/usermodels"
	"github.com/sirupsen/logrus"
)

type QueryUserController struct {
	controller.BaseControllerImpl
}

func (c *QueryUserController) Execute(pl interface{}) (result controller.ResponseController) {
	result = controller.NewResponseController()
	filter := pl.(*usermodels.UserQueryPayload)
	logrus.WithField("pl", filter).Debug("user")
	return
}

func NewQueryUserController() controller.BaseController {
	return &QueryUserController{}
}
