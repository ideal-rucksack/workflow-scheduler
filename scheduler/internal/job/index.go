package job

type Client interface {
	// Type 获取任务类型
	Type() (string, error)

	// Execute 执行任务
	Execute() error
}

const (
	CodeGeneration = "code_generation"
)
