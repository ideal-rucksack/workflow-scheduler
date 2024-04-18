package config

import (
	"github.com/ideal-rucksack/workflow-scheduler/pkg/consotants/cfg/scheduler"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logging"
)

func (p SchedulerProperties) makeSchedulerDefault() {
	makeLoggerDefault(p.Logger)
	makeDatasourceDefault(p.Datasource)
}

func makeLoggerDefault(cfg *logging.Config) {
	level := scheduler.LoggerLevel
	encoding := scheduler.LoggerEncoding
	colors := scheduler.LoggerColors
	timeFormat := scheduler.LoggerTimeFormat
	if cfg == nil {
		cfg = &logging.Config{
			Encoding:   encoding,
			Level:      level,
			Colors:     &colors,
			TimeFormat: timeFormat,
		}
		return
	}

	if cfg.Encoding == "" {
		cfg.Encoding = encoding
	}

	if cfg.Level == "" {
		cfg.Level = level
	}

	if cfg.Colors == nil {
		cfg.Colors = &colors
	}

	if cfg.TimeFormat == "" {
		cfg.TimeFormat = timeFormat
	}
}

func makeDatasourceDefault(cfg *DatasourceConfig) {
	// MySQL
	makeMySQLDefault(cfg.MySQL)
}

func makeMySQLDefault(cfg *MySQLConfig) {
	host := scheduler.MySQLHost
	port := scheduler.MySQLPort
	user := scheduler.MySQLUser
	pass := scheduler.MySQLPass
	sqldb := scheduler.MySQLDB
	maxIdleConn := scheduler.MYSQLMaxIdleConn
	maxOpenConn := scheduler.MYSQLMaxOpenConn
	lifetime := scheduler.MYSQLMaxLifetime
	query := scheduler.MySQLQuery
	if cfg == nil {
		cfg = &MySQLConfig{
			Host:            host,
			Port:            port,
			User:            user,
			Password:        pass,
			Database:        sqldb,
			MaxOpenConn:     maxOpenConn,
			MaxIdleConn:     maxIdleConn,
			ConnMaxLifetime: lifetime,
			Query:           query,
		}
	}

	if cfg.Host == "" {
		cfg.Host = host
	}

	if cfg.Port == 0 {
		cfg.Port = port
	}

	if cfg.User == "" {
		cfg.User = user
	}

	if cfg.Password == "" {
		cfg.Password = pass
	}

	if cfg.Database == "" {
		cfg.Database = sqldb
	}

	if cfg.MaxOpenConn == 0 {
		cfg.MaxOpenConn = maxOpenConn
	}

	if cfg.MaxIdleConn == 0 {
		cfg.MaxIdleConn = maxIdleConn
	}

	if cfg.ConnMaxLifetime == 0 {
		cfg.ConnMaxLifetime = lifetime
	}

	if cfg.Query == "" {
		cfg.Query = query
	}
}
