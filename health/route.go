package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	BuildTime = ""
)

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"buildTime": BuildTime,
	})
}
