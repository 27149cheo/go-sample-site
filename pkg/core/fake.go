package core

import (
	"github.com/gin-gonic/gin"
)

func FakeResponse(c *gin.Context) map[string]interface{} {
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	if raw != "" {
		path = path + "?" + raw
	}

	return map[string]interface{}{
		"method": c.Request.Method,
		"path": path,
		"clientIP": c.ClientIP(),
		"userAgent": c.Request.UserAgent(),
	}
}
