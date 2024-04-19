package cmd

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logging"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/pojo"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo"
	"time"
)

func SetupScheduler(workflowRepo *repo.WorkflowRepo) gocron.Scheduler {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}
	initScheduler(workflowRepo, scheduler)
	defer func(scheduler gocron.Scheduler) {
		err := scheduler.Shutdown()
		if err != nil {
			logging.Logger.Error(err.Error())
		}
	}(scheduler)
	scheduler.Start()
	return scheduler
}

func initScheduler(workflowRepo *repo.WorkflowRepo, scheduler gocron.Scheduler) {
	page, err := workflowRepo.SearchWorkflow(pojo.NewDefaultPageable(1, 10))
	if err != nil {
		logging.Logger.Error(err.Error())
	} else {
		records := page.Records
		if len(records) > 0 {
			for i := range records {
				workflow := records[i]
				if *workflow.Enabled {
					schedule := workflow.Schedule
					workflowJob, err := scheduler.NewJob(
						gocron.CronJob(*schedule, true),
						gocron.NewTask(
							func() {
								logging.Logger.Infof("workflow %s is running", *workflow.Name)
							},
						))
					if err != nil {
						logging.Logger.Fatal(err.Error())
					}
					if workflowJob != nil {
						workerId := workflowJob.ID().ID()
						workflow.WorkerId = &workerId
						now := time.Now()
						workflow.ModifyAt = &now
						_, err = workflowRepo.UpdateWorkflow(workflow)
						if err != nil {
							logging.Logger.Error(err.Error())
						}
					}
				}
			}
		}
	}
}
