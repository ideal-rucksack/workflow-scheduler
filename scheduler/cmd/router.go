package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/config"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/middleware"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/api"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/service"
	"strconv"
)

var (
	jobAPI      *api.JobAPI
	workflowAPI *api.WorkflowAPI
	accountAPI  *api.AccountAPI
)

func setupGin() *gin.Engine {
	engine := gin.Default()
	return engine
}

func setupAPI(engine *gin.Engine) {
	engine.Use(middleware.ErrorMiddleware())

	jobAPI := api.NewJobAPI(service.NewJobService(JobRepository))
	jobAPI.Register(engine)

	workflowAPI := api.NewWorkflowAPI(service.NewWorkflowService(WorkflowRepository))
	workflowAPI.Register(engine)

	accountAPI := api.NewAccountAPI(service.NewAccountService(AccountRepository))
	accountAPI.Register(engine)
}

func setupRestful(cfg config.ServerProperties) error {
	var err error

	engine := setupGin()

	setupAPI(engine)

	if err = engine.Run(":" + strconv.Itoa(cfg.Port)); err != nil {
		return err
	}

	return err
}
