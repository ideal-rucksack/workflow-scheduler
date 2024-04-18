package config

import "github.com/ideal-rucksack/workflow-scheduler/pkg/logging"

type SchedulerProperties struct {
	Datasource *DatasourceConfig `yaml:"datasource"`
	Logger     *logging.Config   `yaml:"logger"`
}

type ConsoleProperties struct {
	Datasource DatasourceConfig `yaml:"datasource"`
	Logger     logging.Config   `yaml:"logger"`
}

type DatasourceConfig struct {
	MySQL *MySQLConfig `yaml:"mysql"`
}

type MySQLConfig struct {
	// Host is the host to connect to the database.
	Host string `mapstructure:"host" yaml:"host"`
	// Port is the port to connect to the database.
	Port int `mapstructure:"port" yaml:"port"`
	// User is the username to connect to the database.
	User string `mapstructure:"user" yaml:"user"`
	// Password is the password to connect to the database.
	Password string `mapstructure:"password" yaml:"password"`
	// Database is the name of the database to connect to.
	Database string `mapstructure:"database" yaml:"database"`
	// MaxOpenConn is the maximum number of open connections to the database.
	MaxOpenConn int `mapstructure:"max_open_conns" yaml:"max_open_conns"`
	// MaxIdleConn is the maximum number of connections in the idle connection pool.
	MaxIdleConn int `mapstructure:"max_idle_conns" yaml:"max_idle_conns"`
	// ConnMaxLifetime is the maximum amount of time a connection may be reused.
	ConnMaxLifetime int `mapstructure:"conn_max_lifetime" yaml:"conn_max_lifetime"`
	// Query is query parameters for connecting to the database
	Query string `mapstructure:"query" yaml:"query"`
}
