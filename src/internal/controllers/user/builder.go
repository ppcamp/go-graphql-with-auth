package user

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/controller"
	postgres "github.com/ppcamp/go-graphql-with-auth/internal/repository"
	"github.com/sirupsen/logrus"
)

type UserControllerBuilder struct {
	handler controller.Handler
	log     *logrus.Entry
}

func NewUserControllerBuilder(storage postgres.Storage) *UserControllerBuilder {
	return &UserControllerBuilder{
		handler: controller.Handler{
			Storage: storage,
		},
		log: logrus.WithField("controller", "UserController"),
	}
}
