package entities

type WorkflowStatus string

const (
	WorkflowSuccess WorkflowStatus = "success"
	WorkflowFailed  WorkflowStatus = "failed"
	WorkflowPending WorkflowStatus = "pending"
	WorkflowStopped WorkflowStatus = "stopped"
	WorkflowTimeout WorkflowStatus = "timeout"
	WorkflowKilled  WorkflowStatus = "killed"
	WorkflowSkipped WorkflowStatus = "skipped"
	WorkflowPaused  WorkflowStatus = "paused"
	WorkflowRunning WorkflowStatus = "running"
)

type Workflow struct {
	GlobalEntity
	Name         *string         `json:"name"`
	Description  *string         `json:"description"`
	Enabled      *bool           `json:"enabled"`
	Status       *WorkflowStatus `json:"status"`        // 状态
	Schedule     *string         `json:"schedule"`      // 调度值: cron则是cron表达式
	ScheduleType *string         `json:"schedule_type"` // 调度类型: cron->定时任务, interval->间隔任务
	CommitAt     *JSONFormatTime `json:"commit_at"`     // 提交时间
}

func (a Workflow) TableName() string {
	return "workflow"
}
