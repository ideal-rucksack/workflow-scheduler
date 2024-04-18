package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/config"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)

func SetupMySQL(cfg config.MySQLConfig) (*sqlx.DB, error) {
	var err error
	url := cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + strconv.Itoa(cfg.Port) + ")/" + cfg.Database + "?" + cfg.Query
	db, err := sqlx.Connect("mysql", url)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetConnMaxLifetime(time.Minute * time.Duration(cfg.ConnMaxLifetime))
	return db, db.Ping()
}
