package api

import "github.com/gin-gonic/gin"

type API interface {
	Register(engine *gin.Engine, middlewares ...gin.HandlerFunc)
}
