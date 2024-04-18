package cmd

import (
	"github.com/ideal-rucksack/workflow-scheduler/pkg/config"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logging"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/setup/datasource"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo"
	"github.com/jmoiron/sqlx"
)

var WorkflowRepository *repo.WorkflowRepo

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
	SetupScheduler(WorkflowRepository)

	// 启动定时任务

	// 启动http服务
	select {}
}

func createRepositories(db *sqlx.DB) {
	WorkflowRepository = repo.NewWorkflowRepo(db)
}
