package code_generation

import (
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logging"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/job"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo/entities"
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
	return nil
}
