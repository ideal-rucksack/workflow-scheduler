package scheduler

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logger"
)

var schedulerInstance gocron.Scheduler

func SetupScheduler() {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		logger.Log.Fatal("创建 %s", err.Error())
	}
	schedulerInstance = scheduler
}

func GetScheduler() gocron.Scheduler {
	return schedulerInstance
}
