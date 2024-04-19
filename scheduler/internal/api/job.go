package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logging"
	job2 "github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/job"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/service"
	"net/http"
	"strconv"
)

type JobAPI struct {
	service *service.JobService
}

func (a JobAPI) Register(engine *gin.Engine) {
	engine.POST("/job/run", a.runJob)
}

func NewJobAPI(service *service.JobService) *JobAPI {
	return &JobAPI{service: service}
}

func (a JobAPI) runJob(context *gin.Context) {
	idStr := context.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
	}
	job, err := a.service.QueryById(id)
	if err != nil {
		logging.Logger.Error(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client, err := job2.GetClient(job)
	if err != nil {
		logging.Logger.Error(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = client.Execute()
	if err != nil {
		logging.Logger.Error(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, job)
}
