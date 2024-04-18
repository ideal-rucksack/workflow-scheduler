package main

import (
	"github.com/ideal-rucksack/workflow-scheduler/cmd/scheduler"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/config"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logging"
)

func main() {
	schedulerCfg, _, err := config.LoadConfig("scheduler")
	if err != nil {
		logging.Logger.Error(err.Error())
	}

	if schedulerCfg != nil {
		scheduler.Execute(*schedulerCfg)
	}
}
