package scheduler

const (
	LoggerLevel      = "info"
	LoggerEncoding   = "console"
	LoggerColors     = true
	LoggerTimeFormat = "2006-01-02 15:04:05"

	MySQLHost = "127.0.0.1"
	MySQLPort = 3306
	MySQLUser = "root"
	MySQLPass = "root"
	MySQLDB   = "workflow_scheduler"
	// MYSQLMaxOpenConn 最大打开连接数 默认10
	MYSQLMaxOpenConn = 10
	// MYSQLMaxIdleConn 最大空闲连接数 默认5
	MYSQLMaxIdleConn = 5
	// MYSQLMaxLifetime 使分配的连接在关闭之前可以重用的最长时间 默认5分钟
	MYSQLMaxLifetime = 5
	MySQLQuery       = "parseTime=true&charset=utf8mb4"

	ServerPort = 5266
)
