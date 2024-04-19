package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo/entities"
	"github.com/jmoiron/sqlx"
)

type JobRepo struct {
	db *sqlx.DB
}

func (r JobRepo) TableName() string {
	return "job"
}

func (r JobRepo) Columns() []string {
	return []string{
		"id",
		"name",
		"description",
		"type",
		"depend_job_id",
		"workflow_id",
		"status",
		"script",
		"create_at",
		"modify_at",
	}
}

func (r JobRepo) Columns2Query() string {
	columns := r.Columns()
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

func (r JobRepo) QueryById(id int64) (*entities.Job, error) {
	var (
		err      error
		job      entities.Job
		querySql = fmt.Sprintf("select %s from %s where id = ?", r.Columns2Query(), r.TableName())
	)
	err = r.db.Get(&job, querySql, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
	}
	return &job, err
}

func NewJobRepo(db *sqlx.DB) *JobRepo {
	return &JobRepo{db: db}
}
