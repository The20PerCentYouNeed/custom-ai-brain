package routes

import (
	"github.com/The20PerCentYouNeed/custom-ai-brain/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handlers.PingHandler)

	return r
}
