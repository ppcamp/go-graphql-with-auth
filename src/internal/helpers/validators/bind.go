package validators

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/utils"
	"github.com/sirupsen/logrus"
)

func ShouldBind(dict map[string]interface{}, obj interface{}) error {
	err := utils.Bind(dict, obj)
	if err != nil {
		logrus.WithError(err).Fatal("It shouldn't throw an error")
	}

	return Validator.Struct(obj)
}
