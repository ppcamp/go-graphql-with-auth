package app

import (
	"time"

	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/controller"
	"github.com/sirupsen/logrus"
)

type QueryPingController struct {
	controller.TransactionControllerImpl
}

func (c *QueryPingController) Execute(pl interface{}) (result controller.ResponseController) {
	result = controller.NewResponseController()

	status := AppStatus{Postgresql: false}
	start := time.Now()

	err := c.Transaction.Ping()
	status.ConnectionDelay = time.Since(start).Microseconds()
	if err != nil {
		logrus.WithError(err).Warn("failed to connect with database")
		result.SetError(err)
		return
	}

	status.Postgresql = true
	logrus.WithField("status", status).Info("query app and it's working")

	result.SetResponse(status)
	return
}

func NewQueryPingController() controller.TransactionController {
	return &QueryPingController{}
}
