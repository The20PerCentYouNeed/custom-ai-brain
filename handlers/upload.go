package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"slices"

	"github.com/The20PerCentYouNeed/custom-ai-brain/models"
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

	buf := make([]byte, 512)
	n, err := src.Read(buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
		return
	}

	contentType := http.DetectContentType(buf)

	types := []string{"application/pdf", "text/plain", "application/vnd.openxmlformats-officedocument.wordprocessingml.document"}
	if !slices.Contains(types, contentType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported file type"})
		return
	}

	var content bytes.Buffer

	content.Write(buf[:n])

	_, err = content.ReadFrom(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
		return
	}

	fileModel := models.File{
		Type:     filepath.Ext(file.Filename),
		Source:   "files",
		Uri:      file.Filename,
		MimeType: contentType,
	}

	var text string

	switch contentType {
	case "application/pdf":
		text, err = fileModel.ExtractTextFromPDF(content)
	case "text/plain":
		text, err = fileModel.ExtractTextFromTXT(content)
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		text, err = fileModel.ExtractTextFromDOCX(content)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported file type"})
		return
	}

	fmt.Println(text)

	// if err := db.DB.Create(&fileModel).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

}
