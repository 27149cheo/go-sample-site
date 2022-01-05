package handler

import (
	"github.com/gin-gonic/gin"

	"go-sample-site/pkg/core/user/service"
	internalhandler "go-sample-site/pkg/internal/handler"
	"go-sample-site/pkg/log"
)

func CreateUser(c *gin.Context) {
	args := &service.User{}
	if err := c.ShouldBindJSON(args); err != nil {
		internalhandler.JSONResponse(c, func() (interface{}, error) { return nil, err})
		return
	}

	internalhandler.CreatedJSONResponse(c, service.CreateUser(args, log.SugaredLogger()))
}

func ListUsers(c *gin.Context) {
	internalhandler.JSONResponse(c, func() (interface{}, error) { return service.ListUsers(log.SugaredLogger())})
}

func DeleteUser(c *gin.Context) {
	internalhandler.DeletedJSONResponse(c, service.DeleteUser(c.Param("id"), log.SugaredLogger()))
}

func GetUser(c *gin.Context) {
	internalhandler.JSONResponse(c, func() (interface{}, error) { return service.GetUser(c.Param("id"), log.SugaredLogger())})
}
