package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"

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

	buf := make([]byte, 512)
	_, err = src.Read(buf)
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

	path := filepath.Join(utils.StoragePath(), "files", file.Filename)
	dst, err := os.Create(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create destination file"})
		return
	}

	if _, err := dst.Write(buf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to write initial bytes"})
		return
	}

	remainingBytes, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read remaining bytes"})
		return
	}

	if _, err := dst.Write(remainingBytes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to write remaining bytes"})
		return
	}

	src.Close()
	dst.Close()

	fileModel := models.File{
		Type:     filepath.Ext(file.Filename),
		Source:   "files",
		Uri:      file.Filename,
		MimeType: contentType,
	}

	var text string

	switch contentType {
	case "application/pdf":
		text, err = fileModel.ExtractTextFromPDF()
	case "text/plain":
		text, err = fileModel.ExtractTextFromTXT()
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		text, err = fileModel.ExtractTextFromDOCX()
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
