package entities

type Job struct {
	GlobalEntity
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	Type        *string `json:"type" db:"type"`
	DependJobId *int64  `json:"depend_job_id" db:"depend_job_id"`
	WorkflowId  *int64  `json:"workflow_id" db:"workflow_id"`
	Status      *string `json:"status" db:"status"` // 状态
	Script      *string `json:"script" db:"script"`
}
