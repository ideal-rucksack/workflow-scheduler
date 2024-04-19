package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/service"
)

type WorkflowAPI struct {
	service *service.WorkflowService
}

func (a WorkflowAPI) Register(engine *gin.Engine) {
	engine.GET("/job/run", getWorkflow)
}

func NewWorkflowAPI(service *service.WorkflowService) *WorkflowAPI {
	return &WorkflowAPI{service: service}
}

func getWorkflow(context *gin.Context) {

}
