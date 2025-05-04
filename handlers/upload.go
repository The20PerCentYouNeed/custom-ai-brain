package handlers

import (
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/The20PerCentYouNeed/custom-ai-brain/db"
	"github.com/The20PerCentYouNeed/custom-ai-brain/models"
	"github.com/The20PerCentYouNeed/custom-ai-brain/utils"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file required"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer src.Close()

	path := filepath.Join(utils.StoragePath(), "files", file.Filename)

	dst, err := os.Create(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create destination file"})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to write file"})
		return
	}

	contentType := filepath.Ext(file.Filename)

	fileModel := models.File{
		Type:     contentType,
		Source:   "files",
		Uri:      file.Filename,
		MimeType: mime.TypeByExtension(contentType),
	}

	var text string

	switch contentType {
	case ".pdf":
		text, err = fileModel.ExtractTextFromPDF()
	case ".txt":
		text, err = fileModel.ExtractTextFromTXT()
	case ".docx":
		text, err = fileModel.ExtractTextFromDOCX()
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported file type"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	document := models.Document{
		File:    &fileModel,
		Content: text,
	}

	err = document.Chunk(300, 50)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&document).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Document uploaded successfully"})
}
