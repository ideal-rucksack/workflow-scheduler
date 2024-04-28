package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/errors"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/api/requests"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/api/response"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/middleware"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo/entities"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/service"
	"net/http"
)

type AccountAPI struct {
	service *service.AccountService
}

func (a AccountAPI) Register(engine *gin.Engine, middlewares ...gin.HandlerFunc) {
	whitelist := engine.Group("/account", middlewares...)
	whitelist.POST("/signin", a.signIn)
	whitelist.POST("/signup", a.signup)
	whitelist.POST("/verify", a.verify)
	whitelist.POST("/refresh_token", a.refreshToken)

	funks := append(middlewares, middleware.BearerAuthorizationMiddleware())
	auth := engine.Group("/account", funks...)
	auth.GET("/signout", a.signOut)
	auth.GET("/current", a.current)
}

func NewAccountAPI(service *service.AccountService) *AccountAPI {
	return &AccountAPI{service: service}
}

func (a AccountAPI) signIn(context *gin.Context) {
	var (
		payload requests.SignIn
		err     error
		result  *response.SignIn
	)

	err = context.ShouldBindJSON(&payload)
	if err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	if err := payload.Validate(); err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	result, err = a.service.SignIn(payload)

	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, result)
}

func (a AccountAPI) signup(context *gin.Context) {
	var (
		payload        requests.Signup
		err            error
		result         *response.Signup
		accountService = a.service
	)

	err = context.ShouldBindJSON(&payload)
	if err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	if err := payload.Validate(); err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	// 注册的用户重写用户状态
	payload.Status = entities.INACTIVE

	result, err = accountService.Signup(payload)

	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, result)
}

func (a AccountAPI) signOut(context *gin.Context) {
	var (
		err            error
		accountService = a.service
	)

	accountId, exists := context.Get("account_id")
	if !exists {
		panic(errors.NewUnauthorizedError("Unauthorized"))
	}

	token, exists := context.Get("token")

	err = accountService.SignOut(accountId.(int64))

	if err != nil {
		panic(err)
	}

	if exists {
		middleware.TokenStoreCache.RemoveToken(token.(string))
	}

	context.JSON(http.StatusOK, true)
}

// verify 验证账户
func (a AccountAPI) verify(context *gin.Context) {
	var (
		payload        requests.Verify
		err            error
		accountService = a.service
		result         *response.Verify
	)

	err = context.ShouldBindJSON(&payload)
	if err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	if err := payload.Validate(); err != nil {
		panic(errors.NewIllegalArgumentError(err.Error()))
	}

	result, err = accountService.Verify(payload)

	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, result)

}

func (a AccountAPI) current(context *gin.Context) {
	var (
		err            error
		accountService = a.service
	)

	accountId, exists := context.Get("account_id")
	if !exists {
		panic(errors.NewUnauthorizedError("Unauthorized"))
	}

	account, err := accountService.Current(accountId.(int64))

	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, account)

}

func (a AccountAPI) refreshToken(context *gin.Context) {
	var (
		err            error
		accountService = a.service
		payload        requests.RefreshToken
		result         *response.RefreshToken
	)

	err = context.ShouldBindJSON(&payload)
	if err != nil {
		panic(errors.NewUnauthorizedError("Unauthorized"))
	}

	result, err = accountService.RefreshToken(payload)

	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, result)

}
