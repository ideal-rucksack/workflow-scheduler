package entities

import "time"

type Workflow struct {
	GlobalEntity
	Name         *string    `json:"name" db:"name"`
	WorkerId     *uint32    `json:"worker_id" db:"worker_id"`
	Description  *string    `json:"description" db:"description"`
	Enabled      *bool      `json:"enabled" db:"enabled"`
	Status       *string    `json:"status" db:"status"`               // 状态
	Schedule     *string    `json:"schedule" db:"schedule"`           // 调度值: cron则是cron表达式
	ScheduleType *string    `json:"schedule_type" db:"schedule_type"` // 调度类型: cron->定时任务, interval->间隔任务
	CommitAt     *time.Time `json:"commit_at" db:"commit_at"`         // 提交时间
}
