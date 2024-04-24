package code_generation

import (
	"github.com/ideal-rucksack/workflow-scheduler/pkg/consotants/cfg"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logging"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/job"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/plugin"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo/entities"
	"os"
)

func init() {
	job.RegisterClientType(job.CodeGeneration, NewCodeGenerationClient)
}

func NewCodeGenerationClient(job *entities.Job) job.Client {
	return CodeGenerationClient{job: job}
}

type CodeGenerationClient struct {
	job *entities.Job
}

func (c CodeGenerationClient) Type() (string, error) {
	return job.CodeGeneration, nil
}

func (c CodeGenerationClient) Execute() error {
	var (
		err      error
		instance plugin.Plugin
	)

	jobInfo := c.job
	if jobInfo != nil {
		logging.Logger.Infof("code generation jobInfo %s is running", *jobInfo.Name)
	}

	loader := plugin.DefaultPluginLoader{Path: os.Getenv(cfg.PluginHome)}
	instance, err = loader.Load("mysql-go-100")
	if err != nil {
		logging.Logger.Errorf("load plugin mysql error %s", err)
	}

	logging.Logger.Infof("plugin %s loaded", instance.Name())
	instance.Executor("-webhook", "http://localhost:5266/webhooks", "-action", "databases")
	//pluginPath := os.Getenv(cfg.PluginHome) + "/mysql/mysql"
	//err := exec.Command(pluginPath, "-webhook", "http://localhost:5266/webhooks", "-action", "databases").Run()
	return nil
}
