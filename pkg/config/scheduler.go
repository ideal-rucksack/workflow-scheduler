package config

import "github.com/ideal-rucksack/workflow-scheduler/pkg/logging"

type SchedulerConfig struct {
	Datasource *DatasourceConfig `yaml:"datasource"`
	Logger     logging.Config    `yaml:"logger"`
}
