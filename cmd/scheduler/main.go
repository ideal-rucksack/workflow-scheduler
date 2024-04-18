package scheduler

import (
	"github.com/ideal-rucksack/workflow-scheduler/pkg/config"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/cmd"
)

func Execute(cfg config.SchedulerProperties) {
	cmd.Start(cfg)
}
