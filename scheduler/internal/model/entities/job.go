package entities

type Job struct {
	GlobalEntity
	Name        *string `json:"name"`
	Description *string `json:"description"`
	DependJobId *int64  `json:"depend_job_id"` // 依赖任务ID
	WorkflowId  *int64  `json:"workflow_id"`   // 工作流ID
	Status      *string `json:"status"`        // 状态
	Script      *string `json:"script"`        // 脚本
}

func (a Job) TableName() string {
	return "job"
}
