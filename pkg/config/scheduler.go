package config

import (
	"github.com/ideal-rucksack/workflow-scheduler/pkg/consotants/cfg/scheduler"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logging"
)

type SchedulerProperties struct {
	Datasource *DatasourceProperties `yaml:"datasource"`
	Logger     *logging.Config       `yaml:"logger"`
	Server     *ServerProperties     `yaml:"server"`
	Plugin     *PluginProperties     `yaml:"plugin"`
}

func (p *SchedulerProperties) make() {
	p.makeLogger()
	p.makeDatasource()
	p.makeServer()
	p.makePlugin()
}

// makeLogger sets the default values for the logger configuration.
func (p *SchedulerProperties) makeLogger() {
	cfg := p.Logger
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
	p.Logger = cfg
}

func (p *SchedulerProperties) makeDatasource() {
	// MySQL
	datasource := p.Datasource
	if datasource == nil {
		datasource = &DatasourceProperties{}
	}
	datasource.makeMySQL()
	p.Datasource = datasource
}

func (p *DatasourceProperties) makeMySQL() {
	cfg := p.MySQL
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
	p.MySQL = cfg
}

func (p *SchedulerProperties) makeServer() {
	cfg := p.Server
	var port = scheduler.ServerPort
	if cfg == nil {
		cfg = &ServerProperties{
			Port: port,
		}
	}

	if cfg.Port == 0 {
		cfg.Port = port
	}
	p.Server = cfg
}

func (p *SchedulerProperties) makePlugin() {
	cfg := p.Plugin
	var home = scheduler.PluginHome
	if cfg == nil {
		cfg = &PluginProperties{
			Home: home,
		}
	}

	if cfg.Home == "" {
		cfg.Home = home
	}
	p.Plugin = cfg
}
