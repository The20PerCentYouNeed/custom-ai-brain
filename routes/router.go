package routes

import (
	"github.com/The20PerCentYouNeed/custom-ai-brain/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	r.GET("/ping", handlers.PingHandler)

	r.POST("/users", handlers.CreateUserHandler)

	r.POST("/ask", handlers.AnswerQuestion)

	r.GET("/documents", handlers.GetDocuments)
	r.GET("/documents/:id", handlers.GetDocument)
	r.POST("/documents", handlers.CreateDocument)
	r.DELETE("/documents/:id", handlers.DestroyDocument)

	r.POST("/documents/upload", handlers.UploadFile)
}
