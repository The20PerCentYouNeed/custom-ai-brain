package handlers

import (
	"net/http"
	"strings"

	"github.com/The20PerCentYouNeed/custom-ai-brain/db"
	"github.com/The20PerCentYouNeed/custom-ai-brain/models"
	"github.com/The20PerCentYouNeed/custom-ai-brain/services/openai"
	"github.com/gin-gonic/gin"
)

type QuestionInput struct {
	Question string `json:"question" binding:"required"`
}

func AnswerQuestion(c *gin.Context) {
	var input QuestionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid 'question' field in JSON"})
		return
	}

	embedding, err := openai.GenerateEmbedding(input.Question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var docs []models.Document

	err = db.DB.Raw(`
		SELECT * FROM documents
		ORDER BY embedding <-> ?
		LIMIT 3
	`, embedding).Scan(&docs).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch documents"})
		return
	}

	var contentBuilder strings.Builder

	for _, doc := range docs {
		contentBuilder.WriteString(doc.Content + "\n\n")
	}

	answer, err := openai.GenerateAnswer(input.Question, contentBuilder.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Answer": answer})
}
