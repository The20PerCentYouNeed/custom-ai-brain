package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/The20PerCentYouNeed/custom-ai-brain/db"
	"github.com/The20PerCentYouNeed/custom-ai-brain/models"
	"github.com/The20PerCentYouNeed/custom-ai-brain/services/openai"
	"github.com/The20PerCentYouNeed/custom-ai-brain/utils"
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

	var results []struct {
		models.Chunk
		models.File
	}

	err = db.DB.Raw(`
			SELECT c.*, f.*
			FROM chunks c
			INNER JOIN documents d ON c.document_id = d.id
			INNER JOIN files f ON f.id = d.file_id
			ORDER BY c.embedding <-> ?
			LIMIT 3
		`, embedding).Scan(&results).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch chunks"})
		return
	}

	var contentBuilder strings.Builder

	for _, result := range results {

		content, err := os.ReadFile(filepath.Join(utils.StoragePath(), "files", result.File.Uri))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
			return
		}

		contentBuilder.WriteString(string(content)[result.Chunk.StartOffset:result.Chunk.EndOffset] + "\n\n")

	}
	fmt.Println("contentBuilder", contentBuilder.String())
	answer, err := openai.GenerateAnswer(input.Question, contentBuilder.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Answer": answer})
}
