package rest

import (
	"fmt"

	articlehandler "go-sample-site/pkg/core/article/handler"
	authorhandler "go-sample-site/pkg/core/author/handler"
	userhandler "go-sample-site/pkg/core/user/handler"
	"go-sample-site/version"

	"github.com/gin-gonic/gin"
)

func (s *engine) injectRouterGroup(router *gin.RouterGroup) {
	{
		router.GET("", func(c *gin.Context) {
			c.String(200, fmt.Sprintf(
				"Version:\t%s\nBuild Number:\t%s\nGit Commit:\t%s",
				version.Version, version.BuildNumber, version.GitCommit,
			))
		})
		router.GET("version", func(c *gin.Context) {
			c.JSON(200, map[string]interface{}{
				"version": version.Version,
				"buildNumber": version.BuildNumber,
				"gitCommit": version.GitCommit,
			})
		})
	}

	for name, r := range map[string]injector{
		"/api/users":     new(userhandler.Router),
		"/api/authors":     new(authorhandler.Router),
		"/api/articles":     new(articlehandler.Router),
	} {
		r.Inject(router.Group(name))
	}
}

type injector interface {
	Inject(router *gin.RouterGroup)
}
