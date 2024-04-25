package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/errors"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logging"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logging.Logger.Error(fmt.Sprintf("Panic: %v", r))
				respondWithError(context, handleError(r))
			}
		}()
		context.Next()
	}
}

func handleError(r interface{}) errors.Errors {
	switch e := r.(type) {
	case errors.UnauthorizedError:
		return e
	case errors.IllegalArgumentError:
		return e
	default:
		// 处理未知的panic
		return errors.NewInternalServerError(fmt.Sprintf("Unknown error: %v", r))
	}
}

func respondWithError(c *gin.Context, err errors.Errors) {
	// 根据错误类型设置HTTP状态码和返回错误信息
	c.AbortWithStatusJSON(err.GetStatus(), gin.H{
		"error": err.GetError(),
	})
}
