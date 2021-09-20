package handler

import (
	"go-sample-site/pkg/core"

	"github.com/gin-gonic/gin"
)

type Router struct{}

func (*Router) Inject(router *gin.RouterGroup) {
	user := router.Group("")
	{
		user.GET("", func(c *gin.Context) {
			c.JSON(200, core.FakeResponse(c))
		})
		user.POST("", func(c *gin.Context) {
			c.JSON(200, core.FakeResponse(c))
		})
		user.GET(":name", func(c *gin.Context) {
			c.JSON(200, core.FakeResponse(c))
		})
		user.PUT(":name", func(c *gin.Context) {
			c.JSON(200, core.FakeResponse(c))
		})
		user.PATCH(":name", func(c *gin.Context) {
			c.JSON(200, core.FakeResponse(c))
		})
		user.DELETE(":name", func(c *gin.Context) {
			c.JSON(200, core.FakeResponse(c))
		})
	}
}
