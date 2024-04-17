package scheduler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron/v2"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/cfg"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/constants"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logger"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/datasource"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/model/entities"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/model/pojo"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repository"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/scheduler"
	"os"
)

func init() {
	cfg.SetupBootstrap("scheduler")
}

func execScheduler() error {
	scheduler.SetupScheduler()
	// 1. 查询数据库所有开启的调度
	workflowRepository := repository.NewWorkflowRepository(datasource.MySQLInstance)
	pageable := pojo.NewDefaultPageable(1, 10)
	search, err := workflowRepository.SearchWorkflow(pageable)
	if err != nil {
		logger.Log.Error("Failed to search workflow %s", err.Error())
		return err
	}
	// 2. 遍历调度，判断是否需要执行
	records := search.Records
	if len(records) > 0 {
		for _, record := range records {
			workflow := record.(entities.Workflow)
			if *workflow.Enabled {
				workflowScheduler := scheduler.GetScheduler()
				schedule := workflow.Schedule
				_, err := workflowScheduler.NewJob(
					gocron.CronJob(*schedule, true),
					gocron.NewTask(
						func() {
							logger.Log.Info("执行调度 %s", *workflow.Name)
						},
					),
				)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func execServer() {
	gin.SetMode(os.Getenv(constants.GinMode))
	engine := gin.Default()
	err := engine.Run(":5266")
	if err != nil {
		logger.Log.Error("Failed to start server %s", err.Error())
	}
}

func exec() {
	logger.SetupLogger("scheduler")
	datasource.SetupMySQL()
	err := execScheduler()
	if err != nil {
		logger.Log.Fatal("Failed to start scheduler %s", err.Error())
		return
	} else {
		go scheduler.GetScheduler().Start()
		defer func(scheduler gocron.Scheduler) {
			err := scheduler.Shutdown()
			if err != nil {
				logger.Log.Error("Failed to stop scheduler %s", err.Error())
			}
		}(scheduler.GetScheduler())
	}
	execServer()
}

func Run() {
	exec()
}
