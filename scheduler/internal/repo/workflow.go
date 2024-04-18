package repo

import (
	"fmt"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/pojo"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/util"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo/entities"
	"github.com/jmoiron/sqlx"
)

type WorkflowRepo struct {
	db *sqlx.DB
}

func (w WorkflowRepo) TableName() string {
	return "workflow"
}

func (w WorkflowRepo) Columns() []string {
	return []string{
		"id",
		"name",
		"worker_id",
		"description",
		"status",
		"enabled",
		"schedule",
		"schedule_type",
		"commit_at",
		"create_at",
		"modify_at",
	}
}

func (w WorkflowRepo) Columns2Query() string {
	columns := w.Columns()
	var columns2Query string
	for i, column := range columns {
		if i == 0 {
			columns2Query = column
		} else {
			columns2Query = fmt.Sprintf("%s, %s", columns2Query, column)
		}
	}
	return columns2Query
}

func NewWorkflowRepo(db *sqlx.DB) *WorkflowRepo {
	return &WorkflowRepo{db: db}
}

func (w WorkflowRepo) SearchWorkflow(pageable pojo.Pageable) (pojo.Page[entities.Workflow], error) {
	db := w.db
	var (
		page      = pojo.Page[entities.Workflow]{Current: pageable.Current(), QueryCount: pageable.QueryCount(), Total: 0, Records: nil}
		totalSql  = fmt.Sprintf("select count(*) from %s", w.TableName())
		pageSql   = fmt.Sprintf("select %s from %s limit ?, ?", w.Columns2Query(), w.TableName())
		workflows []entities.Workflow
		err       error
	)
	limit, offset := util.CalculateLimit(pageable.Current(), pageable.QueryCount())
	err = db.Get(&page.Total, totalSql)
	if err != nil {
		return page, err
	}
	err = db.Select(&workflows, pageSql, limit, offset)
	if err != nil {
		return page, err
	}
	page.Records = workflows
	return page, nil
}

func (w WorkflowRepo) UpdateWorkflow(workflow entities.Workflow) (int64, error) {
	db := w.db
	sql := fmt.Sprintf("update %s set `name`=?, worker_id=?, description=?, status=?, enabled=?, schedule=?, schedule_type=?, commit_at=?, modify_at=? where id=?", w.TableName())
	result, err := db.Exec(sql, workflow.Name, workflow.WorkerId, workflow.Description, workflow.Status, workflow.Enabled, workflow.Schedule, workflow.ScheduleType, workflow.CommitAt, workflow.ModifyAt, workflow.Id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()

}
