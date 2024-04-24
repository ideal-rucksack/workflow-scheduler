package cmd

import (
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/plugin"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/plugin/golang"
)

func init() {
	plugin.Factories["golang"] = golang.NewGolangPlugin
}
