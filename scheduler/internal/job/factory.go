package job

import (
	"fmt"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo/entities"
)

var ClientFactoryMap = map[string]func(job *entities.Job) Client{}

func RegisterClientType(jobType string, factoryFunc func(job *entities.Job) Client) {
	if factoryFunc == nil {
		panic("scheduler: Register client factory is nil")
	}
	ClientFactoryMap[jobType] = factoryFunc
}

func GetClient(job *entities.Job) (Client, error) {
	if job.Type == nil {
		return nil, fmt.Errorf("job type is nil")
	}
	factoryFunc, ok := ClientFactoryMap[*job.Type]
	if !ok {
		return nil, fmt.Errorf("no client registered for job type %s", *job.Type)
	}
	return factoryFunc(job), nil
}
