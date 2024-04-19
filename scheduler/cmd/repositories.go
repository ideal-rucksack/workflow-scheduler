package cmd

import (
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo"
	"github.com/jmoiron/sqlx"
)

var (
	WorkflowRepository *repo.WorkflowRepo
	JobRepository      *repo.JobRepo
)

func createRepositories(db *sqlx.DB) {
	WorkflowRepository = repo.NewWorkflowRepo(db)
	JobRepository = repo.NewJobRepo(db)
}
