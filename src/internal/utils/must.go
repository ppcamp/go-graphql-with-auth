package utils

import (
	"github.com/sirupsen/logrus"
)

func Must(obj interface{}, err error) interface{} {
	if err != nil {
		logrus.WithError(err).Fatal("Failed in some MUST statment")
	}
	return obj
}
