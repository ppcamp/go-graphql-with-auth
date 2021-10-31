package user

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/repository/postgres"
)

type UserControllerBuilder struct {
	// handler controller.Handler
}

func NewUserControllerBuilder(storage postgres.Storage) *UserControllerBuilder {
	return &UserControllerBuilder{
		// handler: controller.Handler{
		// 	Storage: storage,
		// },
	}
}
