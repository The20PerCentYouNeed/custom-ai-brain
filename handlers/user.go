package handlers

import (
	"net/http"

	"github.com/The20PerCentYouNeed/custom-ai-brain/db"
	"github.com/The20PerCentYouNeed/custom-ai-brain/models"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
