package service

import "github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo"

type WorkflowService struct {
	repo *repo.WorkflowRepo
}

func NewWorkflowService(repo *repo.WorkflowRepo) *WorkflowService {
	return &WorkflowService{repo: repo}
}
