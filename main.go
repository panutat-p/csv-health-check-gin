package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/panutat-p/line-remoteinterview-gin/health"
	"github.com/panutat-p/line-remoteinterview-gin/upload"
	"log"
)

const PORT = 8080

var (
	buildTime   = ""
	buildCommit = ""
)

func main() {
	health.BuildTime = buildTime
	health.BuildCommit = buildCommit

	r := gin.Default()
	gin.SetMode(gin.DebugMode) // 🟧 Switch to "release" mode in production

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Origin", "TransactionID"}
	r.Use(cors.New(config))

	r.GET("/", health.CheckHealth)
	r.POST("/upload", upload.CreateFile)

	err := r.Run(fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Println(err)
		log.Println("🟥 Cannot start web server")
		return
	}
}
