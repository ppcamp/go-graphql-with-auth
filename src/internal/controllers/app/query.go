package app

import (
	"time"

	postgres "github.com/ppcamp/go-graphql-with-auth/internal/repository"
)

func (t *AppController) queryPing(tr postgres.Transaction, _ interface{}) (interface{}, error) {
	status := AppStatus{Postgresql: false}
	start := time.Now()
	err := tr.Ping()
	status.ConnectionDelay = time.Since(start).Microseconds()

	if err != nil {
		t.log.WithError(err).Warn("failed to connect with database")
		return nil, err
	}
	status.Postgresql = true

	t.log.WithField("status", status).Info("query app and it's working")

	return status, nil
}
