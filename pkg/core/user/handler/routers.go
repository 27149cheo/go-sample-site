package handler

import (
	"github.com/gin-gonic/gin"
)

type Router struct{}

func (*Router) Inject(router *gin.RouterGroup) {
	connector := router.Group("users")
	{
		connector.POST("", CreateUser)
		connector.GET("", ListUsers)
		connector.GET("/:id", GetUser)
		connector.DELETE("/:id", DeleteUser)
	}
}
