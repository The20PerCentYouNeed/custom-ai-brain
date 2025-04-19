package routes

import (
	"github.com/The20PerCentYouNeed/custom-ai-brain/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handlers.PingHandler)

	r.POST("/users", handlers.CreateUserHandler)

	r.GET("/documents", handlers.GetDocuments)
	r.GET("/documents/:id", handlers.GetDocument)
	r.POST("/documents", handlers.CreateDocument)
	r.DELETE("/documents/:id", handlers.DestroyDocument)

	return r
}
