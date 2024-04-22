package cmd

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/config"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logging"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/setup/datasource"
)

var (
	Scheduler gocron.Scheduler
)

func Start(cfg config.SchedulerProperties) {
	// 初始化日志
	cfg.Logger.CreateLogger()
	// 初始化数据库
	db, err := datasource.SetupMySQL(*cfg.Datasource.MySQL)
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}
	if db == nil {
		logging.Logger.Fatal("mysql is not connected")
	}

	createRepositories(db)

	// 初始化定时任务
	Scheduler = SetupScheduler(WorkflowRepository)
	defer func(scheduler gocron.Scheduler) {
		err := scheduler.Shutdown()
		if err != nil {
			logging.Logger.Error(err.Error())
		}
	}(Scheduler)

	// 启动http服务
	err = setupRestful(*cfg.Server)
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}
}
