package app

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/controller"
	postgres "github.com/ppcamp/go-graphql-with-auth/internal/repository"
	"github.com/sirupsen/logrus"
)

type AppController struct {
	handler controller.Handler
	log     *logrus.Entry
}

func NewAppController(storage postgres.Storage) *AppController {
	return &AppController{
		handler: controller.Handler{
			Storage: storage,
		},
		log: logrus.WithField("controller", "AppController"),
	}
}
