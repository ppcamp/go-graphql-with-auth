package config

var Database = &DatabaseConfig{}

type DatabaseConfig struct {
	Driver         string `json:"RDBMS_DRIVE"`
	Url            string `json:"RDBMS_URL"`
	MaxConnections int    `json:"RDBMS_MAX_CONNECTIONS"`
	MinConnections int    `json:"RDBMS_MIN_CONNECTIONS"`
}
