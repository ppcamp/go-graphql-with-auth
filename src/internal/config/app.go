package config

import "time"

var App = &appConfig{}

type appConfig struct {
	Environment    string        `json:"APP_ENV"`
	Address        string        `json:"APP_ADDRESS"`
	LogLevel       string        `json:"APP_LOG_LEVEL"`
	JWTSecret      string        `json:"APP_JWT_SECRET"`
	JWTExp         time.Duration `json:"APP_JWT_EXP"`
	Migrate        bool          `json:"APP_MIGRATE"`
	ApiDomain      string        `json:"APP_API_DOMAIN"`
	LogPrettyPrint bool          `json:"APP_LOG_PRETTY"`
}

const (
	Development string = "development"
	Staging     string = "staging"
	Production  string = "production"
)
