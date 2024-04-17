package job

type Job interface {

	// Run 运行任务
	Run() error
}
