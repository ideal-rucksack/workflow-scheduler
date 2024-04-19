package service

import (
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo/entities"
)

type JobService struct {
	repo *repo.JobRepo
}

func NewJobService(repo *repo.JobRepo) *JobService {
	return &JobService{repo: repo}
}

func (s JobService) QueryById(id int64) (*entities.Job, error) {
	return s.repo.QueryById(id)
}
