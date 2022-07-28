package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	BuildTime   = ""
	BuildCommit = ""
)

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":      "healthy",
		"buildTime":   BuildTime,
		"buildCommit": BuildCommit,
	})
}
