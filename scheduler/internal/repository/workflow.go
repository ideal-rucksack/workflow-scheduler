package repository

import (
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/model/entities"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/model/pojo"
	"gorm.io/gorm"
)

type WorkflowRepository struct {
	db *gorm.DB
}

func NewWorkflowRepository(db *gorm.DB) *WorkflowRepository {
	return &WorkflowRepository{db: db}
}

func (r WorkflowRepository) SearchWorkflow(pageable pojo.Pageable) (*pojo.Page, error) {
	var (
		workflows []entities.Workflow
		db        = r.db
	)
	result := db.Offset((pageable.Current() - 1) * pageable.QueryCount()).Limit(pageable.QueryCount()).Find(&workflows)
	if result.Error != nil {
		return nil, result.Error
	}

	// 创建一个interface{}类型的切片用于存储转换后的workflows
	records := make([]interface{}, len(workflows))
	for i, workflow := range workflows {
		records[i] = workflow
	}

	return &pojo.Page{
		Current:    pageable.Current(),
		QueryCount: pageable.QueryCount(),
		Records:    records,
	}, nil
}
