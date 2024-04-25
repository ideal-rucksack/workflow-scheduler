package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logging"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/service"
)

type WorkflowAPI struct {
	service *service.WorkflowService
}

func (a WorkflowAPI) Register(engine *gin.Engine, middlewares ...gin.HandlerFunc) {
	engine.POST("/webhooks", a.webhooks)
}

func NewWorkflowAPI(service *service.WorkflowService) *WorkflowAPI {
	return &WorkflowAPI{service: service}
}

func (a WorkflowAPI) getWorkflow(context *gin.Context) {

}

type Response struct {
	Status  int     `json:"status"`
	Error   *string `json:"error"`
	Payload any     `json:"payload"`
}

func (a WorkflowAPI) webhooks(context *gin.Context) {
	var request Response
	_ = context.BindJSON(&request)

	marshal, _ := json.Marshal(request)
	logging.Logger.Infof("webhooks request: %s", string(marshal))
}
