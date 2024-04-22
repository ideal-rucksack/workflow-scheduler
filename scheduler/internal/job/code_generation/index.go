package code_generation

import (
	"github.com/ideal-rucksack/workflow-scheduler/pkg/consotants/cfg"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logging"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/job"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo/entities"
	"os"
	"os/exec"
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
	jobInfo := c.job
	if jobInfo != nil {
		logging.Logger.Infof("code generation jobInfo %s is running", *jobInfo.Name)
	}

	pluginPath := os.Getenv(cfg.PluginHome) + "/mysql/mysql_go_100"
	err := exec.Command(pluginPath, "-webhook", "http://localhost:5266/webhooks", "-action", "databases").Run()
	return err
}
