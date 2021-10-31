package config

import (
	"github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	//#region: App
	&cli.StringFlag{
		Name:        "app_environment",
		Destination: &App.Environment,
		EnvVars:     []string{"APP_ENV"},
		Value:       Development,
	},
	&cli.StringFlag{
		Name:        "app_address",
		Destination: &App.Address,
		EnvVars:     []string{"APP_ADDRESS"},
		Value:       ":8080",
	},
	&cli.StringFlag{
		Name:        "app_log_level",
		Destination: &App.LogLevel,
		EnvVars:     []string{"APP_LOG_LEVEL"},
		Value:       "debug",
	},
	&cli.StringFlag{
		Name:        "app_jwt_secret",
		Destination: &App.JWTSecret,
		EnvVars:     []string{"APP_JWT_SECRET"},
		Value:       "20994458adf248e6a2e2034235e3a0f4",
	},
	&cli.BoolFlag{
		Name:        "app_migrate",
		Destination: &App.Migrate,
		EnvVars:     []string{"APP_MIGRATE"},
		Value:       true,
	},
	&cli.StringFlag{
		Name:        "app_api_domain",
		Destination: &App.ApiDomain,
		EnvVars:     []string{"APP_API_DOMAIN"},
		Value:       "",
	},
	&cli.BoolFlag{
		Name:        "app_log_pretty",
		Destination: &App.LogPrettyPrint,
		EnvVars:     []string{"APP_LOG_PRETTY"},
		Value:       false,
	},
	//#endregion

	//#region: Setup database
	&cli.StringFlag{
		Name:        "rdbms_drive",
		Destination: &Database.Driver,
		EnvVars:     []string{"RDBMS_DRIVE"},
		Value:       "postgres",
	},
	&cli.StringFlag{
		Name:        "rdbms_url",
		Destination: &Database.Url,
		EnvVars:     []string{"RDBMS_URL"},
		Value:       "postgres://gouser:gopsswd@localhost/gousers?sslmode=disable",
	},
	&cli.IntFlag{
		Name:        "rdbms_max_connections",
		Destination: &Database.MaxConnections,
		EnvVars:     []string{"RDBMS_MAX_CONNECTIONS"},
		Value:       10,
	},
	&cli.IntFlag{
		Name:        "rdbms_min_connections",
		Destination: &Database.MinConnections,
		EnvVars:     []string{"RDBMS_MIN_CONNECTIONS"},
		Value:       1,
	},
	//#endregion
}
