package config

import "github.com/sirupsen/logrus"

func Setup() {
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: App.LogPrettyPrint})
	level, err := logrus.ParseLevel(App.LogLevel)

	if err != nil {
		logrus.WithError(err).Fatal("parsing log level")
	}

	logrus.SetLevel(level)

	logrus.WithFields(logrus.Fields{
		"AppConfig":      App,
		"DatabaseConfig": Database,
	}).Info("Environment variables")

}
