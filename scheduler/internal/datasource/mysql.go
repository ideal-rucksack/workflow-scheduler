package datasource

import (
	"fmt"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/constants"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"sync"
)

// MySQLInstance 是数据库连接的单例实例
var MySQLInstance *gorm.DB

// lock 用于确保数据库连接的线程安全
var lock = &sync.Mutex{}

func SetupMySQL() {
	if MySQLInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if MySQLInstance == nil {
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", os.Getenv(constants.DbUser), os.Getenv(constants.DbPassword), os.Getenv(constants.DbHost), os.Getenv(constants.DbPort), os.Getenv(constants.DbName))
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				logger.Log.Fatal("初始化MySQL链接( %s )失败, 因为: %s", dsn, err.Error())
			}
			MySQLInstance = db
		}
	}
}
