package handlers

import (
	"net/http"

	"github.com/The20PerCentYouNeed/custom-ai-brain/db"
	tokenizer "github.com/The20PerCentYouNeed/custom-ai-brain/helpers"
	"github.com/The20PerCentYouNeed/custom-ai-brain/models"
	"github.com/The20PerCentYouNeed/custom-ai-brain/services/openai"
	"github.com/gin-gonic/gin"
)

func GetDocuments(c *gin.Context) {
	var documents []models.Document

	if err := db.DB.Find(&documents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, documents)
}

func GetDocument(c *gin.Context) {
	id := c.Param("id")

	var document models.Document
	if err := db.DB.First(&document, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, document)
}

func CreateDocument(c *gin.Context) {
	var document models.Document
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chunks, err := tokenizer.ChunkByTokens(document.Content, 10, 3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	vectors, err := openai.GenerateEmbeddings(chunks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to generate embeddings",
			"message": err.Error(),
		})
		return
	}

	documentChunks := make([]models.Chunk, len(chunks))
	for i, chunk := range chunks {
		documentChunks[i] = models.Chunk{
			Content:   chunk,
			Embedding: vectors[i],
		}
	}

	document.Chunks = documentChunks

	if err := db.DB.Create(&document).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, document)
}

func DestroyDocument(c *gin.Context) {
	id := c.Param("id")

	var document models.Document
	if err := db.DB.First(&document, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	if err := db.DB.Delete(&document).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
