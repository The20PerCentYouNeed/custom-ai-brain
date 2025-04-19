package handlers

import (
	"net/http"

	"github.com/The20PerCentYouNeed/custom-ai-brain/db"
	"github.com/The20PerCentYouNeed/custom-ai-brain/models"
	"github.com/gin-gonic/gin"
)

func GetDocuments(c *gin.Context) {
	var documents []models.Document

	if err := db.DB.Find(&documents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch documents"})
		return
	}

	c.JSON(http.StatusOK, documents)
}

func GetDocument(c *gin.Context) {
	id := c.Param("id")

	var document models.Document
	if err := db.DB.First(&document, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
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

	if err := db.DB.Create(&document).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create document"})
		return
	}

	c.JSON(http.StatusCreated, document)
}

func DestroyDocument(c *gin.Context) {
	id := c.Param("id")

	var document models.Document
	if err := db.DB.First(&document, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	if err := db.DB.Delete(&document).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
