package config

import (
	"fmt"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/consotants/cfg"
	"github.com/spf13/viper"
	"os"
)

// LoadConfig loads the configuration for the given model
// If model is scheduler, load scheduler configuration
// If the model is console, load the console configuration
// If model is empty, load all configurations
func LoadConfig(models ...string) (*SchedulerProperties, *ConsoleProperties, error) {
	var schedulerProperties SchedulerProperties
	var consoleProperties ConsoleProperties

	if len(models) == 0 {
		return nil, nil, nil
	}

	for i := range models {
		viper.AddConfigPath("conf/")
		viper.SetConfigFile(fmt.Sprintf("conf/%s.yaml", models[i]))
		viper.SetConfigType("yaml")

		if err := viper.ReadInConfig(); err != nil {
			return nil, nil, fmt.Errorf("failed to load %s properties file: %s", models[i], err.Error())
		}

		if models[i] == "scheduler" {
			if err := viper.Unmarshal(&schedulerProperties); err != nil {
				return nil, nil, fmt.Errorf("failed to bind unmarshal %s properties: %s", models[i], err.Error())
			}
			// 未没有设置配置的数据配置默认值
			schedulerProperties.make()
		} else {
			if err := viper.Unmarshal(&consoleProperties); err != nil {
				return nil, nil, fmt.Errorf("failed to bind unmarshal %s properties: %s", models[i], err.Error())
			}
		}
	}
	err := exportOsConfig(schedulerProperties, consoleProperties)
	if err != nil {
		return nil, nil, err
	}

	return &schedulerProperties, &consoleProperties, nil
}

func exportOsConfig(schedule SchedulerProperties, console ConsoleProperties) error {
	err := os.Setenv(cfg.PluginHome, schedule.Plugin.Home)
	return err
}
