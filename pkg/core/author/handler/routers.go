package handler

import (
	"go-sample-site/pkg/core"

	"github.com/gin-gonic/gin"
)

type Router struct{}

func (*Router) Inject(router *gin.RouterGroup) {
	root := router.Group("authors")
	{
		root.GET("", func(c *gin.Context) {
			c.JSON(200, core.FakeResponse(c))
		})
		root.POST("", func(c *gin.Context) {
			c.JSON(200, core.FakeResponse(c))
		})
		root.GET(":name", func(c *gin.Context) {
			c.JSON(200, core.FakeResponse(c))
		})
		root.PUT(":name", func(c *gin.Context) {
			c.JSON(200, core.FakeResponse(c))
		})
		root.PATCH(":name", func(c *gin.Context) {
			c.JSON(200, core.FakeResponse(c))
		})
		root.DELETE(":name", func(c *gin.Context) {
			c.JSON(200, core.FakeResponse(c))
		})
	}
}
